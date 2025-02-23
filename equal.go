package test

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/kylelemons/godebug/diff"
)

func Equal[T comparable](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual == expected, fmt.Sprintf("%v to == %v:\n%s", actual, expected, diff.Diff(spew.Sdump(actual), spew.Sdump(expected))))
}

func NotEqual[T comparable](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual != expected, fmt.Sprintf("%v to == %v", actual, expected))
}

func DeepEqual(t TestingT, actual, expected any) bool {
	t.Helper()

	return assert(t, reflect.DeepEqual(actual, expected), fmt.Sprintf("%v to deep equal %v:\n%s", actual, expected, diff.Diff(spew.Sdump(actual), spew.Sdump(expected))))
}

func NotDeepEqual(t TestingT, actual, expected any) bool {
	t.Helper()

	return assert(t, reflect.DeepEqual(actual, expected), fmt.Sprintf("%v to not deep equal %v", actual, expected))
}
