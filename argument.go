package sqlmock

import "database/sql/driver"

type Argument interface {
	Match(driver.Value) bool
}

func AnyArg() Argument { _ = "STUB: not implemented"; return *new(Argument) }

type anyArgument struct{}

func (a anyArgument) Match(_ driver.Value) bool { _ = "STUB: not implemented"; return false }
