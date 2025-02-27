package test

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/kylelemons/godebug/diff"
)

// Equal asserts that two comparable values are equal.
func Equal[T comparable](t TestingT, actual, expected T) bool {
	t.Helper()

	var zero T
	if _, ok := any(zero).(string); ok {
		return assert(t, actual == expected, fmt.Sprintf("actual to equal expected:\n%s", diff.Diff(any(actual).(string), any(expected).(string))))
	}

	return assert(t, actual == expected, fmt.Sprintf("actual to equal expected:\n%s", diff.Diff(spew.Sdump(actual), spew.Sdump(expected))))
}

// NotEqual asserts that two comparable values are not equal.
func NotEqual[T comparable](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual != expected, fmt.Sprintf("%v to not equal %v", actual, expected))
}

// DeepEqual asserts that two values are deeply equal using reflect.DeepEqual.
func DeepEqual[T any](t TestingT, actual, expected T) bool {
	t.Helper()

	var zero T
	if _, ok := any(zero).(string); ok {
		return assert(t, reflect.DeepEqual(actual, expected), fmt.Sprintf("actual to deep equal expected:\n%s", diff.Diff(any(actual).(string), any(expected).(string))))
	}

	return assert(t, reflect.DeepEqual(actual, expected), fmt.Sprintf("actual to deep equal expected:\n%s", diff.Diff(spew.Sdump(actual), spew.Sdump(expected))))
}

// NotDeepEqual asserts that two values are not deeply equal.
func NotDeepEqual(t TestingT, actual, expected any) bool {
	t.Helper()

	return assert(t, !reflect.DeepEqual(actual, expected), fmt.Sprintf("%v to not deep equal %v", actual, expected))
}
