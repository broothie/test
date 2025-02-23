package test

import "fmt"

func True(t TestingT, actual bool) bool {
	t.Helper()

	return assert(t, actual, fmt.Sprintf("%t to be true", actual))
}

func False(t TestingT, actual bool) bool {
	t.Helper()

	return assert(t, !actual, fmt.Sprintf("%t to be false", actual))
}
