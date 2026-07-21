package sqlmock

import "database/sql/driver"

type SqlMockOption func(*sqlmock) error

func ValueConverterOption(converter driver.ValueConverter) SqlMockOption {
	_ = "STUB: not implemented"
	return *new(SqlMockOption)
}

func QueryMatcherOption(queryMatcher QueryMatcher) SqlMockOption {
	_ = "STUB: not implemented"
	return *new(SqlMockOption)
}

func MonitorPingsOption(monitorPings bool) SqlMockOption {
	_ = "STUB: not implemented"
	return *new(SqlMockOption)
}
