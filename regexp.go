package test

import (
	"fmt"
	"regexp"
)

func RegexpMatch(t TestingT, actual string, expected *regexp.Regexp) bool {
	t.Helper()

	return assert(t, expected.MatchString(actual), fmt.Sprintf("%v to regexp match %v", actual, expected))
}

func NotRegexpMatch(t TestingT, actual string, expected *regexp.Regexp) bool {
	t.Helper()

	return assert(t, !expected.MatchString(actual), fmt.Sprintf("%v to not regexp match %v", actual, expected))
}
