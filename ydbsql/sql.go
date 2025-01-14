package ydbsql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"

	"github.com/yandex-cloud/ydb-go-sdk"
	"github.com/yandex-cloud/ydb-go-sdk/table"
)

var (
	ErrDeprecated          = errors.New("ydbsql: deprecated")
	ErrUnsupported         = errors.New("ydbsql: not supported")
	ErrActiveTransaction   = errors.New("ydbsql: can not begin tx within active tx")
	ErrNoActiveTransaction = errors.New("ydbsql: no active tx to work with")
	ErrSessionBusy         = errors.New("ydbsql: session is busy")
)

var defaultTxControl = table.TxControl(
	table.BeginTx(
		table.WithSerializableReadWrite(),
	),
	table.CommitTx(),
)

// conn is a connection to the ydb.
type conn struct {
	idle bool
	busy bool

	retry   *retryer
	session *table.Session
	pool    *table.SessionPool
	tx      *table.Transaction
	txc     *table.TransactionControl
}

func (c *conn) takeSession(ctx context.Context) bool {
	if !c.idle {
		return true
	}
	if has, _ := c.pool.Take(ctx, c.session); !has {
		return false
	}
	c.idle = false
	return true
}

func (c *conn) putSession(ctx context.Context) {
	err := c.pool.Put(ctx, c.session)
	if err != nil {
		panic(fmt.Sprintf("ydbsql: put session error: %v", err))
	}
	c.idle = true
}

func (c *conn) ResetSession(ctx context.Context) error {
	if c.busy {
		c.pool.PutBusy(ctx, c.session)
		return driver.ErrBadConn
	}
	c.putSession(ctx)
	return nil
}

func (c *conn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if !c.takeSession(ctx) {
		return nil, driver.ErrBadConn
	}
	s, err := c.session.Prepare(ctx, query)
	if err != nil {
		return nil, mapBadSessionError(err)
	}
	return &stmt{
		conn: c,
		stmt: s,
	}, nil
}

// txIsolation maps driver transaction options to ydb transaction option.
// It returns error on unsupported options.
func txIsolation(opts driver.TxOptions) (isolation table.TxOption, err error) {
	level := sql.IsolationLevel(opts.Isolation)
	switch level {
	case sql.LevelDefault,
		sql.LevelSerializable,
		sql.LevelLinearizable:

		return table.WithSerializableReadWrite(), nil

	case sql.LevelReadUncommitted:
		if opts.ReadOnly {
			return table.WithOnlineReadOnly(
				table.WithInconsistentReads(),
			), nil
		}

	case sql.LevelReadCommitted:
		if opts.ReadOnly {
			return table.WithOnlineReadOnly(), nil
		}
	}
	return nil, fmt.Errorf(
		"ydbsql: unsupported transaction options: isolation=%s read_only=%t",
		nameIsolationLevel(level), opts.ReadOnly,
	)
}

func (c *conn) BeginTx(ctx context.Context, opts driver.TxOptions) (tx driver.Tx, err error) {
	if !c.takeSession(ctx) {
		return nil, driver.ErrBadConn
	}
	if c.tx != nil {
		return nil, ErrActiveTransaction
	}
	isolation, err := txIsolation(opts)
	if err != nil {
		return nil, err
	}
	c.tx, err = c.session.BeginTransaction(ctx, table.TxSettings(isolation))
	if err != nil {
		return nil, mapBadSessionError(err)
	}
	c.txc = table.TxControl(table.WithTx(c.tx))
	return c, nil
}

// Rollback implements driver.Tx interface.
// Note that it is called by driver even if a user did not called it.
func (c *conn) Rollback() error {
	if c.tx == nil {
		return ErrNoActiveTransaction
	}

	tx := c.tx
	c.tx = nil
	c.txc = nil

	if c.busy {
		// We don't try to rollback tx here after previous operation was not
		// completed – session is probably still in busy state.
		//
		// Do not return driver.ErrBadConn here – we want this conn's session
		// to be reused after c.ResetSession() call. Bad conn error will force
		// database/sql to close this session without calling ResetSession().
		return ErrSessionBusy
	}

	return tx.Rollback(context.Background())
}

func (c *conn) Commit() error {
	if c.tx == nil {
		return ErrNoActiveTransaction
	}

	tx := c.tx
	c.tx = nil
	c.txc = nil

	if c.busy {
		// We don't try to rollback tx here after previous operation was not
		// completed – session is probably still in busy state.
		//
		// Do not return driver.ErrBadConn here – we want this conn's session
		// to be reused after c.ResetSession() call. Bad conn error will force
		// database/sql to close this session without calling ResetSession().
		return ErrSessionBusy
	}

	return tx.Commit(context.Background())
}

func (c *conn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	txc := c.txc
	if txc == nil {
		txc = defaultTxControl
	}
	_, err := c.exec(ctx, txc, exec{text: query}, params(args))
	if err != nil {
		return nil, err
	}
	return result{}, nil
}

func (c *conn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	txc := c.txc
	if txc == nil {
		txc = defaultTxControl
	}
	res, err := c.exec(ctx, txc, exec{text: query}, params(args))
	if err != nil {
		return nil, err
	}
	res.NextSet()
	return &rows{res: res}, nil
}

func (c *conn) CheckNamedValue(v *driver.NamedValue) error {
	return checkNamedValue(v)
}

func (c *conn) Ping(ctx context.Context) error {
	if !c.takeSession(ctx) {
		return driver.ErrBadConn
	}
	_, err := c.session.KeepAlive(ctx)
	return mapBadSessionError(err)
}

func (c *conn) Close() error {
	if c.busy {
		// NOTE: do not close the session – it is already returned to the pool
		// by PutBusy() call.
		return nil
	}
	ctx := context.Background()
	if !c.takeSession(ctx) {
		return driver.ErrBadConn
	}
	err := c.session.Close(ctx)
	return mapBadSessionError(err)
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return nil, ErrDeprecated
}

func (c *conn) Begin() (driver.Tx, error) {
	return nil, ErrDeprecated
}

// exec is a helper struct for generalization of data query execution.
type exec struct {
	stmt *table.Statement
	text string
}

func (r exec) do(
	ctx context.Context, tx *table.TransactionControl,
	session *table.Session, params *table.QueryParameters,
) (
	*table.Transaction, *table.Result, error,
) {
	if r.stmt != nil {
		return r.stmt.Execute(ctx, tx, params)
	}
	return session.Execute(ctx, tx, r.text, params)
}

func (c *conn) exec(ctx context.Context, tx *table.TransactionControl, exec exec, params *table.QueryParameters) (res *table.Result, err error) {
	if !c.takeSession(ctx) {
		return nil, driver.ErrBadConn
	}
	if c.tx != nil {
		// Under transaction. No need to retry nested calls.
		_, res, err = exec.do(ctx, tx, c.session, params)
	} else {
		// Direct call – retry on errors.
		//
		// NOTE: we do not retrying not found errors here. That is, not found
		// prepared statements are immediately break retry a loop.
		err = c.retry.do(ctx, func(ctx context.Context) (e error) {
			_, res, e = exec.do(ctx, tx, c.session, params)
			return e
		})
	}
	if ydb.IsBusyAfter(err) {
		// Can not use this conn anymore – started operation may not be
		// finished.
		//
		// NOTE: we can not return ErrBadConn right here because we do not
		// know the state of started operation. Instead, we mark this
		// connection busy and it will be closed after ResetSession() call
		// by database/sql.
		c.busy = true
		return nil, err
	}
	return res, mapOpError(err)
}

type TxOperationFunc func(context.Context, *sql.Tx) error

// TxDoer contains options for retrying transactions.
type TxDoer struct {
	DB      *sql.DB
	Options *sql.TxOptions
}

// Do starts a transaction and calls f with it. If f() call returns a retryable
// error, it repeats it accordingly to retry configuration that TxDoer's DB
// driver holds.
//
// Note that callers should mutate state outside of f carefully and keeping in
// mind that f could be called again even if no error returned – transaction
// commitment can be failed:
//
//   var results []int
//   ydbsql.DoTx(ctx, db, TxOperationFunc(func(ctx context.Context, tx *sql.Tx) error {
//       // Reset resulting slice to prevent duplicates when retry occured.
//       results = results[:0]
//
//       rows, err := tx.QueryContext(...)
//       if err != nil {
//           // handle error
//       }
//       for rows.Next() {
//           results = append(results, ...)
//       }
//       return rows.Err()
//   }))
func (td TxDoer) Do(ctx context.Context, f TxOperationFunc) error {
	d := td.DB.Driver().(*Driver)
	return d.c.retry.do(ctx, func(ctx context.Context) error {
		tx, err := td.DB.BeginTx(ctx, td.Options)
		if err != nil {
			return err
		}
		defer tx.Rollback()
		if err := f(ctx, tx); err != nil {
			return err
		}
		return tx.Commit()
	})
}

// DoTx is a shortcut for calling Do(ctx, f) on initialized TxDoer with DB
// field set to given db.
func DoTx(ctx context.Context, db *sql.DB, f TxOperationFunc) error {
	return (TxDoer{DB: db}).Do(ctx, f)
}

var isolationLevelName = [...]string{
	sql.LevelDefault:         "default",
	sql.LevelReadUncommitted: "read_uncommitted",
	sql.LevelReadCommitted:   "read_committed",
	sql.LevelWriteCommitted:  "write_committed",
	sql.LevelRepeatableRead:  "repeatable_read",
	sql.LevelSnapshot:        "snapshot",
	sql.LevelSerializable:    "serializable",
	sql.LevelLinearizable:    "linearizable",
}

func nameIsolationLevel(x sql.IsolationLevel) string {
	if int(x) < len(isolationLevelName) {
		return isolationLevelName[x]
	}
	return "unknown_isolation"
}

type stmt struct {
	conn *conn
	stmt *table.Statement
}

func (s *stmt) NumInput() int {
	return s.stmt.NumInput()
}

func (s *stmt) Close() error {
	return nil
}

func (s stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, ErrDeprecated
}

func (s stmt) Query(args []driver.Value) (driver.Rows, error) {
	return nil, ErrDeprecated
}

func (s *stmt) CheckNamedValue(v *driver.NamedValue) error {
	return checkNamedValue(v)
}

func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	txc := s.conn.txc
	if txc == nil {
		txc = defaultTxControl
	}
	_, err := s.conn.exec(ctx, txc, exec{stmt: s.stmt}, params(args))
	if err != nil {
		return nil, err
	}
	return result{}, nil
}

func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	txc := s.conn.txc
	if txc == nil {
		txc = defaultTxControl
	}
	res, err := s.conn.exec(ctx, txc, exec{stmt: s.stmt}, params(args))
	if err != nil {
		return nil, err
	}
	res.NextSet()
	return &rows{res: res}, nil
}

func checkNamedValue(v *driver.NamedValue) error {
	if v.Name == "" {
		return fmt.Errorf("ydbsql: only named parameters are supported")
	}
	switch x := v.Value.(type) {
	case ydb.Value:
		// OK.

	case valuer:
		// Some ydbsql level types implement valuer interface.
		// Currently it is a date/time types.
		v.Value = x.Value()

	case bool:
		v.Value = ydb.BoolValue(x)
	case int8:
		v.Value = ydb.Int8Value(x)
	case uint8:
		v.Value = ydb.Uint8Value(x)
	case int16:
		v.Value = ydb.Int16Value(x)
	case uint16:
		v.Value = ydb.Uint16Value(x)
	case int32:
		v.Value = ydb.Int32Value(x)
	case uint32:
		v.Value = ydb.Uint32Value(x)
	case int64:
		v.Value = ydb.Int64Value(x)
	case uint64:
		v.Value = ydb.Uint64Value(x)
	case float32:
		v.Value = ydb.FloatValue(x)
	case float64:
		v.Value = ydb.DoubleValue(x)
	case []byte:
		v.Value = ydb.StringValue(x)
	case string:
		v.Value = ydb.UTF8Value(x)
	case [16]byte:
		v.Value = ydb.UUIDValue(x)

	default:
		return fmt.Errorf("ydbsql: unsupported type: %T", x)
	}

	v.Name = "$" + v.Name

	return nil
}

func params(args []driver.NamedValue) *table.QueryParameters {
	if len(args) == 0 {
		return nil
	}
	opts := make([]table.ParameterOption, len(args))
	for i, arg := range args {
		opts[i] = table.ValueParam(
			arg.Name,
			arg.Value.(ydb.Value),
		)
	}
	return table.NewQueryParameters(opts...)
}

type rows struct {
	res *table.Result
}

func (r *rows) Columns() []string {
	var i int
	cs := make([]string, r.res.ColumnCount())
	r.res.Columns(func(m table.Column) {
		cs[i] = m.Name
		i++
	})
	return cs
}

func (r *rows) NextResultSet() error {
	if !r.res.NextSet() {
		return io.EOF
	}
	return nil
}

func (r *rows) HasNextResultSet() bool {
	return r.res.HasNextSet()
}

func (r *rows) Next(dst []driver.Value) error {
	if !r.res.NextRow() {
		return io.EOF
	}
	for i := range dst {
		if !r.res.NextItem() {
			err := r.res.Err()
			if err == nil {
				err = io.ErrUnexpectedEOF
			}
			return err
		}
		if r.res.IsOptional() {
			r.res.Unwrap()
		}
		if r.res.IsDecimal() {
			b, p, s := r.res.UnwrapDecimal()
			dst[i] = Decimal{
				Bytes:     b,
				Precision: p,
				Scale:     s,
			}
		} else {
			dst[i] = r.res.Any()
		}
	}
	return r.res.Err()
}

func (r *rows) Close() error {
	return r.res.Close()
}

type result struct{}

func (r result) LastInsertId() (int64, error) { return 0, ErrUnsupported }
func (r result) RowsAffected() (int64, error) { return 0, ErrUnsupported }

func mapOpError(err error) error {
	switch {
	case ydb.IsOpError(err, ydb.StatusBadSession):
		return driver.ErrBadConn

	case ydb.IsOpError(err, ydb.StatusNotFound):
		// NOTE: if prepared statement is not found, the easy solution is just
		// to drop the sql.Conn (which mapping of table.Session) with all its
		// cached queries.
		//
		// That is, it could be pretty messy to deal with database/sql prepare
		// logic which happens for all sql.Conn instances implicitly.
		return driver.ErrBadConn

	default:
		return err
	}
}

func mapBadSessionError(err error) error {
	if ydb.IsOpError(err, ydb.StatusBadSession) {
		return driver.ErrBadConn
	}
	return err
}
