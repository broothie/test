package test

import (
	"cmp"
	"fmt"
)

// LessThan asserts that a value is less than another value.
func LessThan[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual < expected, fmt.Sprintf("%v to be < %v", actual, expected))
}

// LessThanOrEqual asserts that a value is less than or equal to another value.
func LessThanOrEqual[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual <= expected, fmt.Sprintf("%v to be <= %v", actual, expected))
}

// GreaterThan asserts that a value is greater than another value.
func GreaterThan[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual > expected, fmt.Sprintf("%v to be > %v", actual, expected))
}

// GreaterThanOrEqual asserts that a value is greater than or equal to another value.
func GreaterThanOrEqual[T cmp.Ordered](t TestingT, actual, expected T) bool {
	t.Helper()

	return assert(t, actual >= expected, fmt.Sprintf("%v to be >= %v", actual, expected))
}

// BetweenInclusive asserts that a value is between two values (inclusive).
func BetweenInclusive[T cmp.Ordered](t TestingT, actual, lower, upper T) bool {
	t.Helper()

	return assert(t, actual >= lower && actual <= upper, fmt.Sprintf("%v to be >= %v and <= %v", actual, lower, upper))
}

// BetweenExclusive asserts that a value is between two values (exclusive).
func BetweenExclusive[T cmp.Ordered](t TestingT, actual, lower, upper T) bool {
	t.Helper()

	return assert(t, actual > lower && actual < upper, fmt.Sprintf("%v to be > %v and < %v", actual, lower, upper))
}
