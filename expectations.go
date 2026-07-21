package sqlmock

import (
	"database/sql"
	"database/sql/driver"
	"sync"
	"time"
)

type expectation interface {
	fulfilled() bool
	Lock()
	Unlock()
	String() string
}

type commonExpectation struct {
	sync.Mutex
	triggered bool
	err       error
}

func (e *commonExpectation) fulfilled() bool { _ = "STUB: not implemented"; return false }

type ExpectedClose struct {
	commonExpectation
}

func (e *ExpectedClose) WillReturnError(err error) *ExpectedClose {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedClose) String() string { _ = "STUB: not implemented"; return "" }

type ExpectedBegin struct {
	commonExpectation
	delay  time.Duration
	txOpts *driver.TxOptions
}

func (e *ExpectedBegin) WillReturnError(err error) *ExpectedBegin {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedBegin) String() string { _ = "STUB: not implemented"; return "" }

func (e *ExpectedBegin) WillDelayFor(duration time.Duration) *ExpectedBegin {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedBegin) WithTxOptions(opts sql.TxOptions) *ExpectedBegin {
	_ = "STUB: not implemented"
	return nil
}

type ExpectedCommit struct {
	commonExpectation
}

func (e *ExpectedCommit) WillReturnError(err error) *ExpectedCommit {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedCommit) String() string { _ = "STUB: not implemented"; return "" }

type ExpectedRollback struct {
	commonExpectation
}

func (e *ExpectedRollback) WillReturnError(err error) *ExpectedRollback {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedRollback) String() string { _ = "STUB: not implemented"; return "" }

type ExpectedQuery struct {
	queryBasedExpectation
	rows             driver.Rows
	delay            time.Duration
	rowsMustBeClosed bool
	rowsWereClosed   bool
}

func (e *ExpectedQuery) WithArgs(args ...driver.Value) *ExpectedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedQuery) WithoutArgs() *ExpectedQuery { _ = "STUB: not implemented"; return nil }

func (e *ExpectedQuery) RowsWillBeClosed() *ExpectedQuery { _ = "STUB: not implemented"; return nil }

func (e *ExpectedQuery) WillReturnError(err error) *ExpectedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedQuery) WillDelayFor(duration time.Duration) *ExpectedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedQuery) String() string { _ = "STUB: not implemented"; return "" }

type ExpectedExec struct {
	queryBasedExpectation
	result driver.Result
	delay  time.Duration
}

func (e *ExpectedExec) WithArgs(args ...driver.Value) *ExpectedExec {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedExec) WithoutArgs() *ExpectedExec { _ = "STUB: not implemented"; return nil }

func (e *ExpectedExec) WillReturnError(err error) *ExpectedExec {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedExec) WillDelayFor(duration time.Duration) *ExpectedExec {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedExec) String() string { _ = "STUB: not implemented"; return "" }

func (e *ExpectedExec) WillReturnResult(result driver.Result) *ExpectedExec {
	_ = "STUB: not implemented"
	return nil
}

type ExpectedPrepare struct {
	commonExpectation
	mock         *sqlmock
	expectSQL    string
	statement    driver.Stmt
	closeErr     error
	mustBeClosed bool
	wasClosed    bool
	delay        time.Duration
}

func (e *ExpectedPrepare) WillReturnError(err error) *ExpectedPrepare {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedPrepare) WillReturnCloseError(err error) *ExpectedPrepare {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedPrepare) WillDelayFor(duration time.Duration) *ExpectedPrepare {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedPrepare) WillBeClosed() *ExpectedPrepare { _ = "STUB: not implemented"; return nil }

func (e *ExpectedPrepare) ExpectQuery() *ExpectedQuery { _ = "STUB: not implemented"; return nil }

func (e *ExpectedPrepare) ExpectExec() *ExpectedExec { _ = "STUB: not implemented"; return nil }

func (e *ExpectedPrepare) String() string { _ = "STUB: not implemented"; return "" }

type queryBasedExpectation struct {
	commonExpectation
	expectSQL string
	converter driver.ValueConverter
	args      []driver.Value
	noArgs    bool
}

type ExpectedPing struct {
	commonExpectation
	delay time.Duration
}

func (e *ExpectedPing) WillDelayFor(duration time.Duration) *ExpectedPing {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedPing) WillReturnError(err error) *ExpectedPing {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectedPing) String() string { _ = "STUB: not implemented"; return "" }
