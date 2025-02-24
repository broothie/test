package test

import "fmt"

// NotNil asserts that a value is not nil.
func NotNil(t TestingT, actual any) bool {
	t.Helper()

	return assert(t, actual != nil, fmt.Sprintf("%v to not be nil", actual))
}
