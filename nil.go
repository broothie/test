package test

import (
	"fmt"
	"reflect"
)

// Nil asserts that a value is nil.
func Nil(t TestingT, actual any) bool {
	t.Helper()

	if actual == nil {
		return assert(t, true, fmt.Sprintf("%v to be nil", actual))
	}

	v := reflect.ValueOf(actual)
	isNil := v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func

	return assert(t, isNil && v.IsNil(), fmt.Sprintf("%v to be nil", actual))
}

// NotNil asserts that a value is not nil.
func NotNil(t TestingT, actual any) bool {
	t.Helper()

	if actual == nil {
		return assert(t, false, fmt.Sprintf("%v to not be nil", actual))
	}

	v := reflect.ValueOf(actual)
	isNil := v.Kind() == reflect.Ptr ||
		v.Kind() == reflect.Interface ||
		v.Kind() == reflect.Slice ||
		v.Kind() == reflect.Map ||
		v.Kind() == reflect.Chan ||
		v.Kind() == reflect.Func

	return assert(t, !isNil || !v.IsNil(), fmt.Sprintf("%v to not be nil", actual))
}
