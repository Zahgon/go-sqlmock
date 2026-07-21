//go:build go1.8
// +build go1.8

package sqlmock

import (
	"context"
	"database/sql/driver"
	"errors"
)

type Sqlmock interface {
	SqlmockCommon

	NewRowsWithColumnDefinition(columns ...*Column) *Rows

	NewColumn(name string) *Column
}

var ErrCancelled = errors.New("canceling query due to user request")

func (c *sqlmock) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (c *sqlmock) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (c *sqlmock) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

func (c *sqlmock) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (c *sqlmock) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) ping() (*ExpectedPing, error) { _ = "STUB: not implemented"; return nil, nil }

func (stmt *statement) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (stmt *statement) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (c *sqlmock) ExpectPing() *ExpectedPing { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) Query(query string, args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (c *sqlmock) query(query string, args []driver.NamedValue) (*ExpectedQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *sqlmock) Exec(query string, args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (c *sqlmock) exec(query string, args []driver.NamedValue) (*ExpectedExec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *sqlmock) NewRowsWithColumnDefinition(columns ...*Column) *Rows {
	_ = "STUB: not implemented"
	return nil
}

func (c *sqlmock) NewColumn(name string) *Column { _ = "STUB: not implemented"; return nil }
