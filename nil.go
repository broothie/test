package test

import "fmt"

func Nil(t TestingT, actual any) bool {
	t.Helper()

	return assert(t, actual == nil, fmt.Sprintf("%v to be nil", actual))
}

func NotNil(t TestingT, actual any) bool {
	t.Helper()

	return assert(t, actual != nil, fmt.Sprintf("%v to not be nil", actual))
}
