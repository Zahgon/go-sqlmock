//go:build go1.8
// +build go1.8

package sqlmock

import (
	"database/sql/driver"
)

func (e *ExpectedQuery) WillReturnRows(rows ...*Rows) *ExpectedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (e *queryBasedExpectation) argsMatches(args []driver.NamedValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *queryBasedExpectation) attemptArgMatch(args []driver.NamedValue) (err error) {
	_ = "STUB: not implemented"
	return nil
}
