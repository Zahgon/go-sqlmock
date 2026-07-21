package sqlmock

import (
	"database/sql"
	"database/sql/driver"
)

type SqlmockCommon interface {
	ExpectClose() *ExpectedClose

	ExpectationsWereMet() error

	ExpectPrepare(expectedSQL string) *ExpectedPrepare

	ExpectQuery(expectedSQL string) *ExpectedQuery

	ExpectExec(expectedSQL string) *ExpectedExec

	ExpectBegin() *ExpectedBegin

	ExpectCommit() *ExpectedCommit

	ExpectRollback() *ExpectedRollback

	ExpectPing() *ExpectedPing

	MatchExpectationsInOrder(bool)

	NewRows(columns []string) *Rows
}

type sqlmock struct {
	ordered      bool
	dsn          string
	opened       int
	drv          *mockDriver
	converter    driver.ValueConverter
	queryMatcher QueryMatcher
	monitorPings bool

	expected []expectation
}

func (c *sqlmock) open(options []SqlMockOption) (*sql.DB, Sqlmock, error) {
	_ = "STUB: not implemented"
	return nil, *new(Sqlmock), nil
}

func (c *sqlmock) ExpectClose() *ExpectedClose { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) MatchExpectationsInOrder(b bool) { _ = "STUB: not implemented"; return }

func (c *sqlmock) Close() error { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) ExpectationsWereMet() error { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

func (c *sqlmock) begin(opts driver.TxOptions) (*ExpectedBegin, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *sqlmock) ExpectBegin() *ExpectedBegin { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) ExpectExec(expectedSQL string) *ExpectedExec {
	_ = "STUB: not implemented"
	return nil
}

func (c *sqlmock) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (c *sqlmock) prepare(query string) (*ExpectedPrepare, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *sqlmock) ExpectPrepare(expectedSQL string) *ExpectedPrepare {
	_ = "STUB: not implemented"
	return nil
}

func (c *sqlmock) ExpectQuery(expectedSQL string) *ExpectedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (c *sqlmock) ExpectCommit() *ExpectedCommit { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) ExpectRollback() *ExpectedRollback { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) Commit() error { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) Rollback() error { _ = "STUB: not implemented"; return nil }

func (c *sqlmock) NewRows(columns []string) *Rows { _ = "STUB: not implemented"; return nil }
