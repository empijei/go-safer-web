package safesql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

var (
	// ErrConnDone is https://pkg.go.dev/database/sql#ErrConnDone
	ErrConnDone = sql.ErrConnDone
	// ErrNoRows is https://pkg.go.dev/database/sql#ErrNoRows
	ErrNoRows = sql.ErrNoRows
	// ErrTxDone is https://pkg.go.dev/database/sql#ErrTxDone
	ErrTxDone = sql.ErrTxDone
)

// Drivers is https://pkg.go.dev/sql#Drivers
func Drivers() []string { return sql.Drivers() }

// func is https://pkg.go.dev/database/sql#func
func Register(name string, driver driver.Driver) { sql.Register(name, driver) }

// Forwarded safe types.
type (
	// ColumnType is https://pkg.go.dev/sql#ColumnType
	ColumnType = sql.ColumnType
	// DBStats is https://pkg.go.dev/sql#DBStats
	DBStats = sql.DBStats
	// IsolationLevel is https://pkg.go.dev/sql#IsolationLevel
	IsolationLevel = sql.IsolationLevel
	// NamedArg is https://pkg.go.dev/sql#NamedArg
	NamedArg = sql.NamedArg
	// NullBool is https://pkg.go.dev/sql#NullBool
	NullBool = sql.NullBool
	// NullFloat64 is https://pkg.go.dev/sql#NullFloat64
	NullFloat64 = sql.NullFloat64
	// NullInt32 is https://pkg.go.dev/sql#NullInt32
	NullInt32 = sql.NullInt32
	// NullInt64 is https://pkg.go.dev/sql#NullInt64
	NullInt64 = sql.NullInt64
	// NullString is https://pkg.go.dev/sql#NullString
	NullString = sql.NullString
	// NullTime is https://pkg.go.dev/sql#NullTime
	NullTime = sql.NullTime
	// Out is https://pkg.go.dev/sql#Out
	Out = sql.Out
	// RawBytes is https://pkg.go.dev/sql#RawBytes
	RawBytes = sql.RawBytes
	// Result is https://pkg.go.dev/sql#Result
	Result = sql.Result
	// Row is https://pkg.go.dev/sql#Row
	Row = sql.Row
	// Rows is https://pkg.go.dev/sql#Rows
	Rows = sql.Rows
	// Scanner is https://pkg.go.dev/sql#Scanner
	Scanner = sql.Scanner
	// Stmt is https://pkg.go.dev/sql#Stmt
	Stmt = sql.Stmt
	// TxOptions is https://pkg.go.dev/sql#TxOptions
	TxOptions = sql.TxOptions
)

// Conn is a tiny wrapper for https://pkg.go.dev/database/sql#Conn
// The Raw method has been removed for security reasons.
type Conn struct {
	c *sql.Conn
}

// BeginTx is a tiny wrapper for https://pkg.go.dev/database/sql#Conn.BeginTx
func (c *Conn) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error) {
	t, err := c.c.BeginTx(ctx, opts)
	return &Tx{t}, err
}

// Close is a tiny wrapper for https://pkg.go.dev/database/sql#Conn.Close
func (c *Conn) Close() error {
	return c.c.Close()
}

// ExecContext is a tiny wrapper for https://pkg.go.dev/database/sql#Conn.ExecContext
func (c *Conn) ExecContext(ctx context.Context, query String, args ...any) (Result, error) {
	return c.c.ExecContext(ctx, query.s, args...)
}

// PingContext is a tiny wrapper for https://pkg.go.dev/database/sql#Conn.PingContext
func (c *Conn) PingContext(ctx context.Context) error {
	return c.c.PingContext(ctx)
}

// PrepareContext is a tiny wrapper for https://pkg.go.dev/database/sql#Conn.PrepareContext
func (c *Conn) PrepareContext(ctx context.Context, query String) (*Stmt, error) {
	return c.c.PrepareContext(ctx, query.s)
}

// QueryContext is a tiny wrapper for https://pkg.go.dev/database/sql#Conn.QueryContext
func (c *Conn) QueryContext(ctx context.Context, query String, args ...any) (*Rows, error) {
	return c.c.QueryContext(ctx, query.s, args...)
}

// QueryRowContext is a tiny wrapper for https://pkg.go.dev/database/sql#Conn.QueryRowContext
func (c *Conn) QueryRowContext(ctx context.Context, query String, args ...any) *Row {
	return c.c.QueryRowContext(ctx, query.s, args...)
}

// DB is a tiny wrapper for https://pkg.go.dev/database/sql#DB
// The Driver method has been removed for security reasons.
type DB struct {
	db *sql.DB
}

// Open is a tiny wrapper for https://pkg.go.dev/database/sql#Open
func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	return &DB{db}, err
}

// OpenDB is a tiny wrapper for https://pkg.go.dev/database/sql#OpenDB
func OpenDB(c driver.Connector) *DB {
	return &DB{sql.OpenDB(c)}
}

// Begin is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Begin
func (db *DB) Begin() (*Tx, error) {
	t, err := db.db.Begin()
	return &Tx{t}, err
}

// BeginTx is a tiny wrapper for https://pkg.go.dev/database/sql#DB.BeginTx
func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error) {
	t, err := db.db.BeginTx(ctx, opts)
	return &Tx{t}, err
}

// Close is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Close
func (db *DB) Close() error {
	return db.db.Close()
}

// Conn is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Conn
func (db *DB) Conn(ctx context.Context) (*Conn, error) {
	c, err := db.db.Conn(ctx)
	return &Conn{c}, err
}

// Exec is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Exec
func (db *DB) Exec(query String, args ...any) (Result, error) {
	return db.db.Exec(query.s, args...)
}

// ExecContext is a tiny wrapper for https://pkg.go.dev/database/sql#DB.ExecContext
func (db *DB) ExecContext(ctx context.Context, query String, args ...any) (Result, error) {
	return db.db.ExecContext(ctx, query.s, args...)
}

// Ping is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Ping
func (db *DB) Ping() error {
	return db.db.Ping()
}

// PingContext is a tiny wrapper for https://pkg.go.dev/database/sql#DB.PingContext
func (db *DB) PingContext(ctx context.Context) error {
	return db.db.PingContext(ctx)
}

// Prepare is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Prepare
func (db *DB) Prepare(query String) (*Stmt, error) {
	return db.db.Prepare(query.s)
}

// PrepareContext is a tiny wrapper for https://pkg.go.dev/database/sql#DB.PrepareContext
func (db *DB) PrepareContext(ctx context.Context, query String) (*Stmt, error) {
	return db.db.PrepareContext(ctx, query.s)
}

// Query is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Query
func (db *DB) Query(query String, args ...any) (*Rows, error) {
	return db.db.Query(query.s, args...)
}

// QueryContext is a tiny wrapper for https://pkg.go.dev/database/sql#DB.QueryContext
func (db *DB) QueryContext(ctx context.Context, query String, args ...any) (*Rows, error) {
	return db.db.QueryContext(ctx, query.s, args...)
}

// QueryRow is a tiny wrapper for https://pkg.go.dev/database/sql#DB.QueryRow
func (db *DB) QueryRow(query String, args ...any) *Row {
	return db.db.QueryRow(query.s, args...)
}

// QueryRowContext is a tiny wrapper for https://pkg.go.dev/database/sql#DB.QueryRowContext
func (db *DB) QueryRowContext(ctx context.Context, query String, args ...any) *Row {
	return db.db.QueryRowContext(ctx, query.s, args...)
}

// SetConnMaxIdleTime is a tiny wrapper for https://pkg.go.dev/database/sql#DB.SetConnMaxIdleTime
func (db *DB) SetConnMaxIdleTime(d time.Duration) {
	db.db.SetConnMaxIdleTime(d)
}

// SetConnMaxLifetime is a tiny wrapper for https://pkg.go.dev/database/sql#DB.SetConnMaxLifetime
func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.db.SetConnMaxLifetime(d)
}

// SetMaxIdleConns is a tiny wrapper for https://pkg.go.dev/database/sql#DB.SetMaxIdleConns
func (db *DB) SetMaxIdleConns(n int) {
	db.db.SetMaxIdleConns(n)
}

// SetMaxOpenConns is a tiny wrapper for https://pkg.go.dev/database/sql#DB.SetMaxOpenConns
func (db *DB) SetMaxOpenConns(n int) {
	db.db.SetMaxOpenConns(n)
}

// Stats is a tiny wrapper for https://pkg.go.dev/database/sql#DB.Stats
func (db *DB) Stats() DBStats {
	return db.db.Stats()
}

// Tx is a tiny wrapper for https://pkg.go.dev/database/sql#Tx
type Tx struct {
	tx *sql.Tx
}

// Commit is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.Commit
func (tx *Tx) Commit() error {
	return tx.tx.Commit()
}

// Exec is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.Exec
func (tx *Tx) Exec(query String, args ...any) (Result, error) {
	return tx.tx.Exec(query.s, args...)
}

// ExecContext is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.ExecContext
func (tx *Tx) ExecContext(ctx context.Context, query String, args ...any) (Result, error) {
	return tx.tx.ExecContext(ctx, query.s, args...)
}

// Prepare is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.Prepare
func (tx *Tx) Prepare(query String) (*Stmt, error) {
	return tx.tx.Prepare(query.s)
}

// PrepareContext is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.PrepareContext
func (tx *Tx) PrepareContext(ctx context.Context, query String) (*Stmt, error) {
	return tx.tx.PrepareContext(ctx, query.s)
}

// Query is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.Query
func (tx *Tx) Query(query String, args ...any) (*Rows, error) {
	return tx.tx.Query(query.s, args...)
}

// QueryContext is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.QueryContext
func (tx *Tx) QueryContext(ctx context.Context, query String, args ...any) (*Rows, error) {
	return tx.tx.QueryContext(ctx, query.s, args...)
}

// QueryRow is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.QueryRow
func (tx *Tx) QueryRow(query String, args ...any) *Row {
	return tx.tx.QueryRow(query.s, args...)
}

// QueryRowContext is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.QueryRowContext
func (tx *Tx) QueryRowContext(ctx context.Context, query String, args ...any) *Row {
	return tx.tx.QueryRowContext(ctx, query.s, args...)
}

// Rollback is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.Rollback
func (tx *Tx) Rollback() error {
	return tx.tx.Rollback()
}

// Stmt is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.Stmt
func (tx *Tx) Stmt(stmt *Stmt) *Stmt {
	return tx.tx.Stmt(stmt)
}

// StmtContext is a tiny wrapper for https://pkg.go.dev/database/sql#Tx.StmtContext
func (tx *Tx) StmtContext(ctx context.Context, stmt *Stmt) *Stmt {
	return tx.tx.StmtContext(ctx, stmt)
}
