package match

import (
	"fmt"
	"maps"
	"slices"
)

func Contains[T comparable](expected T) MatcherFunc[[]T] {
	return func(subject []T) (string, bool) {
		return fmt.Sprintf("contain %v", expected), slices.Contains(subject, expected)
	}
}

func ContainsKey[K comparable, V any](expected K) MatcherFunc[map[K]V] {
	return func(subject map[K]V) (string, bool) {
		_, ok := subject[expected]

		return fmt.Sprintf("contain key %v", expected), ok
	}
}

func ContainsValue[K, V comparable](expected V) MatcherFunc[map[K]V] {
	return func(subject map[K]V) (string, bool) {
		return fmt.Sprintf("contain value %v", expected), slices.Contains(slices.Collect(maps.Values(subject)), expected)
	}
}

func ContainsEntry[K, V comparable](expectedKey K, expectedValue V) MatcherFunc[map[K]V] {
	return func(subject map[K]V) (string, bool) {
		return fmt.Sprintf("contain entry %v", map[K]V{expectedKey: expectedValue}), slices.ContainsFunc(slices.Collect(maps.Keys(subject)), func(key K) bool {
			return key == expectedKey && subject[key] == expectedValue
		})
	}
}
