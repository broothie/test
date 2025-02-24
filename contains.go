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
