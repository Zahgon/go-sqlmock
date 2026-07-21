package sqlmock

type statement struct {
	conn  *sqlmock
	ex    *ExpectedPrepare
	query string
}

func (stmt *statement) Close() error { _ = "STUB: not implemented"; return nil }

func (stmt *statement) NumInput() int { _ = "STUB: not implemented"; return 0 }
