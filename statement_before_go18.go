//go:build !go1.8
// +build !go1.8

package sqlmock

import (
	"database/sql/driver"
)

func (stmt *statement) Exec(args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (stmt *statement) Query(args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}
