package test

import (
	"errors"
	"fmt"
)

// ErrorMessageIs asserts that an error's message matches the expected string.
func ErrorMessageIs(t TestingT, actual error, expected string) bool {
	t.Helper()

	return assert(t, actual.Error() == expected, fmt.Sprintf("%v to have message %q", actual, expected))
}

// NotErrorMessageIs asserts that an error's message does not match the expected string.
func NotErrorMessageIs(t TestingT, actual error, expected string) bool {
	t.Helper()

	return assert(t, actual.Error() != expected, fmt.Sprintf("%v to not have message %q", actual, expected))
}

// ErrorIs asserts that an error matches another error using errors.Is.
func ErrorIs(t TestingT, actual, expected error) bool {
	t.Helper()

	return assert(t, errors.Is(actual, expected), fmt.Sprintf("%v to wrap %v", actual, expected))
}

// NotErrorIs asserts that an error does not match another error using errors.Is.
func NotErrorIs(t TestingT, actual, expected error) bool {
	t.Helper()

	return assert(t, !errors.Is(actual, expected), fmt.Sprintf("%v to not wrap %v", actual, expected))
}

// NoError asserts that an error is nil.
func NoError(t TestingT, err error) bool {
	t.Helper()

	return assert(t, err == nil, fmt.Sprintf("%v to be nil", err))
}

// MustNoError fails the test immediately if the error is not nil.
func MustNoError(t TestingT, err error) {
	t.Helper()

	if !NoError(t, err) {
		t.FailNow()
	}
}

// Must runs a function that returns a value and error, fails immediately if error is not nil.
func Must[T any](t TestingT, f func() (T, error)) T {
	t.Helper()

	result, err := f()
	MustNoError(t, err)
	return result
}
