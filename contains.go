package test

import (
	"fmt"
	"maps"
	"slices"
)

func Contains[T comparable, S ~[]T](t TestingT, actual S, expected T) bool {
	t.Helper()

	return assert(t, slices.Contains(actual, expected), fmt.Sprintf("%v to contain %v", actual, expected))
}

func NotContains[T comparable, S ~[]T](t TestingT, actual S, expected T) bool {
	t.Helper()

	return assert(t, !slices.Contains(actual, expected), fmt.Sprintf("%v to not contain %v", actual, expected))
}

func ContainsKey[K comparable, V any, M ~map[K]V](t TestingT, actual M, expected K) bool {
	t.Helper()

	return assert(t, slices.Contains(slices.Collect(maps.Keys(actual)), expected), fmt.Sprintf("%v to contain %v", actual, expected))
}

func NotContainsKey[K comparable, V any, M ~map[K]V](t TestingT, actual M, expected K) bool {
	t.Helper()

	return assert(t, !slices.Contains(slices.Collect(maps.Keys(actual)), expected), fmt.Sprintf("%v to not contain %v", actual, expected))
}

func ContainsValue[K, V comparable, M ~map[K]V](t TestingT, actual M, expected V) bool {
	t.Helper()

	return assert(t, slices.Contains(slices.Collect(maps.Values(actual)), expected), fmt.Sprintf("%v to contain %v", actual, expected))
}

func NotContainsValue[K, V comparable, M ~map[K]V](t TestingT, actual M, expected V) bool {
	t.Helper()

	return assert(t, !slices.Contains(slices.Collect(maps.Values(actual)), expected), fmt.Sprintf("%v to not contain %v", actual, expected))
}

func ContainsEntry[K, V comparable, M ~map[K]V](t TestingT, actual M, expectedKey K, expectedValue V) bool {
	t.Helper()

	assertion := slices.ContainsFunc(slices.Collect(maps.Keys(actual)), func(key K) bool { return key == expectedKey && actual[key] == expectedValue })
	return assert(t, assertion, fmt.Sprintf("%v to contain %v", actual, map[K]V{expectedKey: expectedValue}))
}

func NotContainsEntry[K, V comparable, M ~map[K]V](t TestingT, actual M, expectedKey K, expectedValue V) bool {
	t.Helper()

	assertion := !slices.ContainsFunc(slices.Collect(maps.Keys(actual)), func(key K) bool { return key == expectedKey && actual[key] == expectedValue })
	return assert(t, assertion, fmt.Sprintf("%v to not contain %v", actual, map[K]V{expectedKey: expectedValue}))
}
