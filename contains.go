package test

import (
	"fmt"
	"maps"
	"slices"
)

// Contains asserts that a slice contains an expected element.
func Contains[T comparable, S ~[]T](t TestingT, actual S, expected T) bool {
	t.Helper()

	return assert(t, slices.Contains(actual, expected), fmt.Sprintf("%v to contain %v", actual, expected))
}

// NotContains asserts that a slice does not contain an expected element.
func NotContains[T comparable, S ~[]T](t TestingT, actual S, expected T) bool {
	t.Helper()

	return assert(t, !slices.Contains(actual, expected), fmt.Sprintf("%v to not contain %v", actual, expected))
}

// ContainsKey asserts that a map contains an expected key.
func ContainsKey[K comparable, V any, M ~map[K]V](t TestingT, actual M, expected K) bool {
	t.Helper()

	return assert(t, slices.Contains(slices.Collect(maps.Keys(actual)), expected), fmt.Sprintf("%v to contain %v", actual, expected))
}

// NotContainsKey asserts that a map does not contain an expected key.
func NotContainsKey[K comparable, V any, M ~map[K]V](t TestingT, actual M, expected K) bool {
	t.Helper()

	return assert(t, !slices.Contains(slices.Collect(maps.Keys(actual)), expected), fmt.Sprintf("%v to not contain %v", actual, expected))
}

// ContainsValue asserts that a map contains an expected value.
func ContainsValue[K, V comparable, M ~map[K]V](t TestingT, actual M, expected V) bool {
	t.Helper()

	return assert(t, slices.Contains(slices.Collect(maps.Values(actual)), expected), fmt.Sprintf("%v to contain %v", actual, expected))
}

// NotContainsValue asserts that a map does not contain an expected value.
func NotContainsValue[K, V comparable, M ~map[K]V](t TestingT, actual M, expected V) bool {
	t.Helper()

	return assert(t, !slices.Contains(slices.Collect(maps.Values(actual)), expected), fmt.Sprintf("%v to not contain %v", actual, expected))
}

// ContainsEntry asserts that a map contains an expected key-value pair.
func ContainsEntry[K, V comparable, M ~map[K]V](t TestingT, actual M, expectedKey K, expectedValue V) bool {
	t.Helper()

	assertion := slices.ContainsFunc(slices.Collect(maps.Keys(actual)), func(key K) bool { return key == expectedKey && actual[key] == expectedValue })
	return assert(t, assertion, fmt.Sprintf("%v to contain %v", actual, map[K]V{expectedKey: expectedValue}))
}

// NotContainsEntry asserts that a map does not contain an expected key-value pair.
func NotContainsEntry[K, V comparable, M ~map[K]V](t TestingT, actual M, expectedKey K, expectedValue V) bool {
	t.Helper()

	assertion := !slices.ContainsFunc(slices.Collect(maps.Keys(actual)), func(key K) bool { return key == expectedKey && actual[key] == expectedValue })
	return assert(t, assertion, fmt.Sprintf("%v to not contain %v", actual, map[K]V{expectedKey: expectedValue}))
}

// ContainedIn asserts that an element is contained in a slice.
func ContainedIn[T comparable, S ~[]T](t TestingT, actual T, expected S) bool {
	t.Helper()

	return assert(t, slices.Contains(expected, actual), fmt.Sprintf("%v to be contained in %v", actual, expected))
}

// NotContainedIn asserts that an element is not contained in a slice.
func NotContainedIn[T comparable, S ~[]T](t TestingT, actual T, expected S) bool {
	t.Helper()

	return assert(t, !slices.Contains(expected, actual), fmt.Sprintf("%v to not be contained in %v", actual, expected))
}

// ContainedInKeys asserts that a key is contained in a map.
func ContainedInKeys[K comparable, V any, M ~map[K]V](t TestingT, actual K, expected M) bool {
	t.Helper()

	return assert(t, slices.Contains(slices.Collect(maps.Keys(expected)), actual), fmt.Sprintf("%v to be contained in keys of %v", actual, expected))
}

// NotContainedInKeys asserts that a key is not contained in a map.
func NotContainedInKeys[K comparable, V any, M ~map[K]V](t TestingT, actual K, expected M) bool {
	t.Helper()

	return assert(t, !slices.Contains(slices.Collect(maps.Keys(expected)), actual), fmt.Sprintf("%v to not be contained in keys of %v", actual, expected))
}

// ContainedInValues asserts that a value is contained in a map.
func ContainedInValues[K, V comparable, M ~map[K]V](t TestingT, actual V, expected M) bool {
	t.Helper()

	return assert(t, slices.Contains(slices.Collect(maps.Values(expected)), actual), fmt.Sprintf("%v to be contained in values of %v", actual, expected))
}

// NotContainedInValues asserts that a value is not contained in a map.
func NotContainedInValues[K, V comparable, M ~map[K]V](t TestingT, actual V, expected M) bool {
	t.Helper()

	return assert(t, !slices.Contains(slices.Collect(maps.Values(expected)), actual), fmt.Sprintf("%v to not be contained in values of %v", actual, expected))
}

// ContainedInEntries asserts that a key-value pair is contained in a map.
func ContainedInEntries[K, V comparable, M ~map[K]V](t TestingT, actualKey K, actualValue V, expected M) bool {
	t.Helper()

	return assert(t, slices.ContainsFunc(slices.Collect(maps.Keys(expected)), func(key K) bool { return key == actualKey && expected[key] == actualValue }), fmt.Sprintf("%v to contain entries %v", expected, map[K]V{actualKey: actualValue}))
}

// NotContainedInEntries asserts that a key-value pair is not contained in a map.
func NotContainedInEntries[K, V comparable, M ~map[K]V](t TestingT, actualKey K, actualValue V, expected M) bool {
	t.Helper()

	return assert(t, !slices.ContainsFunc(slices.Collect(maps.Keys(expected)), func(key K) bool { return key == actualKey && expected[key] == actualValue }), fmt.Sprintf("%v to not contain entries %v", expected, map[K]V{actualKey: actualValue}))
}
