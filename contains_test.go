package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestContains(t *testing.T) {
	t.Run("when the subject contains the expected value", shouldPass(func(t test.TestingT) {
		test.Contains(t, []int{1, 2, 3}, 2)
	}))

	t.Run("when the subject does not contain the expected value", shouldFail(func(t test.TestingT) {
		test.Contains(t, []int{1, 2, 3}, 4)
	}))
}

func TestNotContains(t *testing.T) {
	t.Run("when the subject does not contain the expected value", shouldPass(func(t test.TestingT) {
		test.NotContains(t, []int{1, 2, 3}, 4)
	}))

	t.Run("when the subject contains the expected value", shouldFail(func(t test.TestingT) {
		test.NotContains(t, []int{1, 2, 3}, 2)
	}))
}

func TestContainsKey(t *testing.T) {
	t.Run("when the map contains the expected key", shouldPass(func(t test.TestingT) {
		test.ContainsKey(t, map[string]int{"foo": 1, "bar": 2}, "foo")
	}))

	t.Run("when the map does not contain the expected key", shouldFail(func(t test.TestingT) {
		test.ContainsKey(t, map[string]int{"foo": 1, "bar": 2}, "baz")
	}))
}

func TestNotContainsKey(t *testing.T) {
	t.Run("when the map does not contain the expected key", shouldPass(func(t test.TestingT) {
		test.NotContainsKey(t, map[string]int{"foo": 1, "bar": 2}, "baz")
	}))

	t.Run("when the map contains the expected key", shouldFail(func(t test.TestingT) {
		test.NotContainsKey(t, map[string]int{"foo": 1, "bar": 2}, "foo")
	}))
}

func TestContainsValue(t *testing.T) {
	t.Run("when the map contains the expected value", shouldPass(func(t test.TestingT) {
		test.ContainsValue(t, map[string]int{"foo": 1, "bar": 2}, 1)
	}))

	t.Run("when the map does not contain the expected value", shouldFail(func(t test.TestingT) {
		test.ContainsValue(t, map[string]int{"foo": 1, "bar": 2}, 3)
	}))
}

func TestNotContainsValue(t *testing.T) {
	t.Run("when the map does not contain the expected value", shouldPass(func(t test.TestingT) {
		test.NotContainsValue(t, map[string]int{"foo": 1, "bar": 2}, 3)
	}))

	t.Run("when the map contains the expected value", shouldFail(func(t test.TestingT) {
		test.NotContainsValue(t, map[string]int{"foo": 1, "bar": 2}, 1)
	}))
}

func TestContainsEntry(t *testing.T) {
	t.Run("when the map contains the expected key-value pair", shouldPass(func(t test.TestingT) {
		test.ContainsEntry(t, map[string]int{"foo": 1, "bar": 2}, "foo", 1)
	}))

	t.Run("when the map does not contain the expected key", shouldFail(func(t test.TestingT) {
		test.ContainsEntry(t, map[string]int{"foo": 1, "bar": 2}, "baz", 1)
	}))

	t.Run("when the map contains the key but with different value", shouldFail(func(t test.TestingT) {
		test.ContainsEntry(t, map[string]int{"foo": 1, "bar": 2}, "foo", 2)
	}))
}

func TestNotContainsEntry(t *testing.T) {
	t.Run("when the map does not contain the expected key-value pair", shouldPass(func(t test.TestingT) {
		test.NotContainsEntry(t, map[string]int{"foo": 1, "bar": 2}, "baz", 3)
	}))

	t.Run("when the map contains the key but with different value", shouldPass(func(t test.TestingT) {
		test.NotContainsEntry(t, map[string]int{"foo": 1, "bar": 2}, "foo", 2)
	}))

	t.Run("when the map contains the expected key-value pair", shouldFail(func(t test.TestingT) {
		test.NotContainsEntry(t, map[string]int{"foo": 1, "bar": 2}, "foo", 1)
	}))
}
