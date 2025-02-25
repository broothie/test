package test

import "fmt"

func SliceLen[T any, S ~[]T](t TestingT, actual S, expected int) bool {
	t.Helper()

	return assert(t, len(actual) == expected, fmt.Sprintf("%v to be of length %d", actual, expected))
}

func SliceEmpty[T any, S ~[]T](t TestingT, actual S) bool {
	t.Helper()

	return assert(t, len(actual) == 0, fmt.Sprintf("%v to be empty", actual))
}

func NotSliceEmpty[T any, S ~[]T](t TestingT, actual S) bool {
	t.Helper()

	return assert(t, len(actual) != 0, fmt.Sprintf("%v to not be empty", actual))
}

func MapLen[K comparable, V any, M ~map[K]V](t TestingT, actual M, expected int) bool {
	t.Helper()

	return assert(t, len(actual) == expected, fmt.Sprintf("%v to be of length %d", actual, expected))
}

func MapEmpty[K comparable, V any, M ~map[K]V](t TestingT, actual M) bool {
	t.Helper()

	return assert(t, len(actual) == 0, fmt.Sprintf("%v to be empty", actual))
}

func NotMapEmpty[K comparable, V any, M ~map[K]V](t TestingT, actual M) bool {
	t.Helper()

	return assert(t, len(actual) != 0, fmt.Sprintf("%v to not be empty", actual))
}
