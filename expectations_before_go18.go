//go:build !go1.8
// +build !go1.8

package sqlmock

func (e *ExpectedQuery) WillReturnRows(rows *Rows) *ExpectedQuery {
	_ = "STUB: not implemented"
	return nil
}

func (e *queryBasedExpectation) argsMatches(args []namedValue) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *queryBasedExpectation) attemptArgMatch(args []namedValue) (err error) {
	_ = "STUB: not implemented"
	return nil
}
