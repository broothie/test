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

	return assert(t, actual == expected, fmt.Sprintf("%v to == %v:\n%s", actual, expected, diff.Diff(spew.Sdump(actual), spew.Sdump(expected))))
}

// NotEqual asserts that two comparable values are not equal.
func NotEqual[T comparable](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual != expected, fmt.Sprintf("%v to == %v", actual, expected))
}

// DeepEqual asserts that two values are deeply equal using reflect.DeepEqual.
func DeepEqual(t TestingT, actual, expected any) bool {
	t.Helper()

	return assert(t, reflect.DeepEqual(actual, expected), fmt.Sprintf("%v to deep equal %v:\n%s", actual, expected, diff.Diff(spew.Sdump(actual), spew.Sdump(expected))))
}

// NotDeepEqual asserts that two values are not deeply equal.
func NotDeepEqual(t TestingT, actual, expected any) bool {
	t.Helper()

	return assert(t, !reflect.DeepEqual(actual, expected), fmt.Sprintf("%v to not deep equal %v", actual, expected))
}
