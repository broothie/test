package test

import (
	"fmt"
	"reflect"

	"github.com/broothie/test/match"
)

func Expect[Subject any](t TestingT, subject Subject, match match.Matcher[Subject]) bool {
	t.Helper()

	to, isMatch := match.Match(subject)
	if !isMatch {
		if isFunc(subject) {
			t.Error(fmt.Sprintf("expected subject to %s", to))
		} else {
			t.Error(fmt.Sprintf("expected %v to %s", subject, to))
		}
	}

	return isMatch
}

func isFunc(a any) bool {
	return reflect.TypeOf(a).Kind() == reflect.Func
}
