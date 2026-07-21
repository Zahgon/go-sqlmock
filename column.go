package sqlmock

import "reflect"

type Column struct {
	name       string
	dbType     string
	nullable   bool
	nullableOk bool
	length     int64
	lengthOk   bool
	precision  int64
	scale      int64
	psOk       bool
	scanType   reflect.Type
}

func (c *Column) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Column) DbType() string { _ = "STUB: not implemented"; return "" }

func (c *Column) IsNullable() (bool, bool) { _ = "STUB: not implemented"; return false, false }

func (c *Column) Length() (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func (c *Column) PrecisionScale() (int64, int64, bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

func (c *Column) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func NewColumn(name string) *Column { _ = "STUB: not implemented"; return nil }

func (c *Column) Nullable(nullable bool) *Column { _ = "STUB: not implemented"; return nil }

func (c *Column) OfType(dbType string, sampleValue interface{}) *Column {
	_ = "STUB: not implemented"
	return nil
}

func (c *Column) WithLength(length int64) *Column { _ = "STUB: not implemented"; return nil }

func (c *Column) WithPrecisionAndScale(precision, scale int64) *Column {
	_ = "STUB: not implemented"
	return nil
}
