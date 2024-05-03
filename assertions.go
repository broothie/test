package test

import (
	"errors"
)

const errorfFormat = "got %+v, want %+v"

type TestingT interface {
	Helper()
	Errorf(string, ...any)
	FailNow()
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

func assertf[T any](t TestingT, assertion bool, got, want T) bool {
	t.Helper()

	if !assertion {
		t.Errorf(errorfFormat, got, want)
	}

	return assertion
}
