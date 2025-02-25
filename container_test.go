package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestSliceLen(t *testing.T) {
	t.Run("when slice length matches expected", shouldPass(func(t test.TestingT) {
		test.SliceLen(t, []int{1, 2, 3}, 3)
	}))

	t.Run("when slice length does not match expected", shouldFail(func(t test.TestingT) {
		test.SliceLen(t, []int{1, 2}, 3)
	}))
}

func TestSliceEmpty(t *testing.T) {
	t.Run("when slice is empty", shouldPass(func(t test.TestingT) {
		test.SliceEmpty(t, []int{})
	}))

	t.Run("when slice is not empty", shouldFail(func(t test.TestingT) {
		test.SliceEmpty(t, []int{1})
	}))
}

func TestNotSliceEmpty(t *testing.T) {
	t.Run("when slice is not empty", shouldPass(func(t test.TestingT) {
		test.NotSliceEmpty(t, []int{1})
	}))

	t.Run("when slice is empty", shouldFail(func(t test.TestingT) {
		test.NotSliceEmpty(t, []int{})
	}))
}

func TestMapLen(t *testing.T) {
	t.Run("when map length matches expected", shouldPass(func(t test.TestingT) {
		test.MapLen(t, map[string]int{"a": 1, "b": 2}, 2)
	}))

	t.Run("when map length does not match expected", shouldFail(func(t test.TestingT) {
		test.MapLen(t, map[string]int{"a": 1}, 2)
	}))
}

func TestMapEmpty(t *testing.T) {
	t.Run("when map is empty", shouldPass(func(t test.TestingT) {
		test.MapEmpty(t, map[string]int{})
	}))

	t.Run("when map is not empty", shouldFail(func(t test.TestingT) {
		test.MapEmpty(t, map[string]int{"a": 1})
	}))
}

func TestNotMapEmpty(t *testing.T) {
	t.Run("when map is not empty", shouldPass(func(t test.TestingT) {
		test.NotMapEmpty(t, map[string]int{"a": 1})
	}))

	t.Run("when map is empty", shouldFail(func(t test.TestingT) {
		test.NotMapEmpty(t, map[string]int{})
	}))
}
