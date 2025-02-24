package test

import (
	"errors"
	"fmt"
)

func ErrorMessageIs(t TestingT, actual error, expected string) bool {
	t.Helper()

	return assert(t, actual.Error() == expected, fmt.Sprintf("%v to have message %q", actual, expected))
}

func NotErrorMessageIs(t TestingT, actual error, expected string) bool {
	t.Helper()

	return assert(t, actual.Error() != expected, fmt.Sprintf("%v to not have message %q", actual, expected))
}

func ErrorIs(t TestingT, actual, expected error) bool {
	t.Helper()

	return assert(t, errors.Is(actual, expected), fmt.Sprintf("%v to wrap %v", actual, expected))
}

func NotErrorIs(t TestingT, actual, expected error) bool {
	t.Helper()

	return assert(t, !errors.Is(actual, expected), fmt.Sprintf("%v to not wrap %v", actual, expected))
}

func MustNoError(t TestingT, err error) {
	t.Helper()

	if !Equal(t, err, nil) {
		t.FailNow()
	}
}

func Must[T any](t TestingT, f func() (T, error)) T {
	t.Helper()

	result, err := f()
	MustNoError(t, err)
	return result
}
