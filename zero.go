package test

import (
	"fmt"
	"reflect"
)

// Zero asserts that a value equals its zero value.
// Works with any type by using reflection.
func Zero(t TestingT, actual any) bool {
	t.Helper()

	return assert(t, reflect.ValueOf(actual).IsZero(), fmt.Sprintf("%v to be zero value", actual))
}

// NotZero asserts that a value does not equal its zero value.
// Works with any type by using reflection.
func NotZero(t TestingT, actual any) bool {
	t.Helper()

	return assert(t, !reflect.ValueOf(actual).IsZero(), fmt.Sprintf("%v to not be zero value", actual))
}
