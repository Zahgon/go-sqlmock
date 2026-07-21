package sqlmock

import (
	"database/sql/driver"
)

type result struct {
	insertID     int64
	rowsAffected int64
	err          error
}

func NewResult(lastInsertID int64, rowsAffected int64) driver.Result {
	_ = "STUB: not implemented"
	return *new(driver.Result)
}

func NewErrorResult(err error) driver.Result { _ = "STUB: not implemented"; return *new(driver.Result) }

func (r *result) LastInsertId() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *result) RowsAffected() (int64, error) { _ = "STUB: not implemented"; return 0, nil }
