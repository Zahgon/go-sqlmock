package sqlmock

import (
	"database/sql/driver"
	"strings"
)

const invalidate = "☠☠☠ MEMORY OVERWRITTEN ☠☠☠ "

var CSVColumnParser = func(s string) interface{} {
	switch {
	case strings.ToLower(s) == "null":
		return nil
	}
	return []byte(s)
}

type rowSets struct {
	sets []*Rows
	pos  int
	ex   *ExpectedQuery
	raw  [][]byte
}

func (rs *rowSets) Columns() []string { _ = "STUB: not implemented"; return nil }

func (rs *rowSets) Close() error { _ = "STUB: not implemented"; return nil }

func (rs *rowSets) Next(dest []driver.Value) error { _ = "STUB: not implemented"; return nil }

func (rs *rowSets) String() string { _ = "STUB: not implemented"; return "" }

func (rs *rowSets) empty() bool { _ = "STUB: not implemented"; return false }

func rawBytes(col driver.Value) (_ []byte, ok bool) { _ = "STUB: not implemented"; return nil, false }

func (rs *rowSets) invalidateRaw() { _ = "STUB: not implemented"; return }

type Rows struct {
	converter driver.ValueConverter
	cols      []string
	def       []*Column
	rows      [][]driver.Value
	pos       int
	nextErr   map[int]error
	closeErr  error
}

func NewRows(columns []string) *Rows { _ = "STUB: not implemented"; return nil }

func (r *Rows) CloseError(err error) *Rows { _ = "STUB: not implemented"; return nil }

func (r *Rows) RowError(row int, err error) *Rows { _ = "STUB: not implemented"; return nil }

func (r *Rows) AddRow(values ...driver.Value) *Rows { _ = "STUB: not implemented"; return nil }

func (r *Rows) AddRows(values ...[]driver.Value) *Rows { _ = "STUB: not implemented"; return nil }

func (r *Rows) FromCSVString(s string) *Rows { _ = "STUB: not implemented"; return nil }
