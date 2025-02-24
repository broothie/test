package test

import (
	"fmt"
	"regexp"
)

// RegexpMatch asserts that a string matches a regular expression.
func RegexpMatch(t TestingT, actual string, expected *regexp.Regexp) bool {
	t.Helper()

	return assert(t, expected.MatchString(actual), fmt.Sprintf("%v to regexp match %v", actual, expected))
}

// NotRegexpMatch asserts that a string does not match a regular expression.
func NotRegexpMatch(t TestingT, actual string, expected *regexp.Regexp) bool {
	t.Helper()

	return assert(t, !expected.MatchString(actual), fmt.Sprintf("%v to not regexp match %v", actual, expected))
}
