//go:build go1.8 && !go1.9
// +build go1.8,!go1.9

package sqlmock

import "database/sql/driver"

func (c *sqlmock) CheckNamedValue(nv *driver.NamedValue) (err error) {
	_ = "STUB: not implemented"
	return nil
}
