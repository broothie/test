package test

import (
	"cmp"
	"fmt"
)

func LessThan[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual < expected, fmt.Sprintf("%v to be < %v", actual, expected))
}

func LessThanOrEqual[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual <= expected, fmt.Sprintf("%v to be <= %v", actual, expected))
}

func GreaterThan[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual > expected, fmt.Sprintf("%v to be > %v", actual, expected))
}

func GreaterThanOrEqual[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual >= expected, fmt.Sprintf("%v to be >= %v", actual, expected))
}

func BetweenInclusive[T cmp.Ordered](t TestingT, actual, lower, upper T) bool {
	t.Helper()

	return assert(t, actual >= lower && actual <= upper, fmt.Sprintf("%v to be >= %v and <= %v", actual, lower, upper))
}

func BetweenExclusive[T cmp.Ordered](t TestingT, actual, lower, upper T) bool {
	t.Helper()

	return assert(t, actual > lower && actual < upper, fmt.Sprintf("%v to be > %v and < %v", actual, lower, upper))
}
