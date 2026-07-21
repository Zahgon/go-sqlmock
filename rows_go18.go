//go:build go1.8
// +build go1.8

package sqlmock

import (
	"reflect"
)

func (rs *rowSets) HasNextResultSet() bool { _ = "STUB: not implemented"; return false }

func (rs *rowSets) NextResultSet() error { _ = "STUB: not implemented"; return nil }

type rowSetsWithDefinition struct {
	*rowSets
}

func (rs *rowSetsWithDefinition) ColumnTypeDatabaseTypeName(index int) string {
	_ = "STUB: not implemented"
	return ""
}

func (rs *rowSetsWithDefinition) ColumnTypeLength(index int) (length int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (rs *rowSetsWithDefinition) ColumnTypeNullable(index int) (nullable, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (rs *rowSetsWithDefinition) ColumnTypePrecisionScale(index int) (precision, scale int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

func (rs *rowSetsWithDefinition) ColumnTypeScanType(index int) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (rs *rowSetsWithDefinition) getDefinition(index int) *Column {
	_ = "STUB: not implemented"
	return nil
}

func NewRowsWithColumnDefinition(columns ...*Column) *Rows { _ = "STUB: not implemented"; return nil }
