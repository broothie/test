package test

import "fmt"

// True asserts that the actual value is true.
func True(t TestingT, actual bool) bool {
	t.Helper()

	return assert(t, actual, fmt.Sprintf("%t to be true", actual))
}

// False asserts that the actual value is false.
func False(t TestingT, actual bool) bool {
	t.Helper()

	return assert(t, !actual, fmt.Sprintf("%t to be false", actual))
}
