package sqlmock

import (
	"database/sql"
	"database/sql/driver"
	"sync"
)

var pool *mockDriver

func init() {
	pool = &mockDriver{
		conns: make(map[string]*sqlmock),
	}
	sql.Register("sqlmock", pool)
}

type mockDriver struct {
	sync.Mutex
	counter int
	conns   map[string]*sqlmock
}

func (d *mockDriver) Open(dsn string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func New(options ...SqlMockOption) (*sql.DB, Sqlmock, error) {
	_ = "STUB: not implemented"
	return nil, *new(Sqlmock), nil
}

func NewWithDSN(dsn string, options ...SqlMockOption) (*sql.DB, Sqlmock, error) {
	_ = "STUB: not implemented"
	return nil, *new(Sqlmock), nil
}
