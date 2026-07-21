package sqlmock

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("\\s+")

func stripQuery(q string) (s string) { _ = "STUB: not implemented"; return "" }

type QueryMatcher interface {
	Match(expectedSQL, actualSQL string) error
}

type QueryMatcherFunc func(expectedSQL, actualSQL string) error

func (f QueryMatcherFunc) Match(expectedSQL, actualSQL string) error {
	_ = "STUB: not implemented"
	return nil
}

var QueryMatcherRegexp QueryMatcher = QueryMatcherFunc(func(expectedSQL, actualSQL string) error {
	expect := stripQuery(expectedSQL)
	actual := stripQuery(actualSQL)
	if actual != "" && expect == "" {
		return fmt.Errorf("expectedSQL can't be empty")
	}
	re, err := regexp.Compile(expect)
	if err != nil {
		return err
	}
	if !re.MatchString(actual) {
		return fmt.Errorf(`could not match actual sql: "%s" with expected regexp "%s"`, actual, re.String())
	}
	return nil
})

var QueryMatcherEqual QueryMatcher = QueryMatcherFunc(func(expectedSQL, actualSQL string) error {
	expect := stripQuery(expectedSQL)
	actual := stripQuery(actualSQL)
	if actual != expect {
		return fmt.Errorf(`actual sql: "%s" does not equal to expected "%s"`, actual, expect)
	}
	return nil
})
