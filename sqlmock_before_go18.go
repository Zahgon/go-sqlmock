//go:build !go1.8
// +build !go1.8

package sqlmock

import (
	"database/sql/driver"
)

type Sqlmock interface {
	SqlmockCommon
}

type namedValue struct {
	Name    string
	Ordinal int
	Value   driver.Value
}

func (c *sqlmock) ExpectPing() *ExpectedPing { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) Query(query string, args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (c *sqlmock) query(query string, args []namedValue) (*ExpectedQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *sqlmock) Exec(query string, args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (c *sqlmock) exec(query string, args []namedValue) (*ExpectedExec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
