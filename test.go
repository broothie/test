package test

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/kylelemons/godebug/diff"
)

func DeepEqual(t TestingT, got, want any) bool {
	t.Helper()

	return assertf(t, reflect.DeepEqual(got, want), got, want)
}

func Equal[T comparable](t TestingT, got, want T) bool {
	t.Helper()

	return assertf(t, got == want, got, want)
}

func True(t TestingT, got bool) bool {
	t.Helper()

	return Equal(t, got, true)
}

func False(t TestingT, got bool) bool {
	t.Helper()

	return Equal(t, got, false)
}

func ErrorIs(t TestingT, got, want error) bool {
	t.Helper()

	return assertf(t, errors.Is(got, want), got, want)
}

func NoError(t TestingT, err error) bool {
	t.Helper()

	return ErrorIs(t, err, nil)
}

func MustNoError(t TestingT, err error) {
	t.Helper()

	if !NoError(t, err) {
		t.FailNow()
	}
}

func Must[T any](t TestingT, f func() (T, error)) T {
	t.Helper()

	result, err := f()
	MustNoError(t, err)
	return result
}

func assertf(t TestingT, assertion bool, got, want any) bool {
	t.Helper()

	if !assertion {
		t.Error(fmt.Sprintf("\n%s", diff.Diff(spew.Sdump(got), spew.Sdump(want))))
	}

	return assertion
}
