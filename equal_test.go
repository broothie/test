package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestEqual(t *testing.T) {
	t.Run("when values are equal", shouldPass(func(t test.TestingT) {
		test.Equal(t, 1, 1)
	}))

	t.Run("when values are not equal", shouldFail(func(t test.TestingT) {
		test.Equal(t, 1, 2)
	}))

	t.Run("when strings are equal", shouldPass(func(t test.TestingT) {
		test.Equal(t, "foo", "foo")
	}))
}

func TestNotEqual(t *testing.T) {
	t.Run("when values are not equal", shouldPass(func(t test.TestingT) {
		test.NotEqual(t, 1, 2)
	}))

	t.Run("when values are equal", shouldFail(func(t test.TestingT) {
		test.NotEqual(t, 1, 1)
	}))
}

func TestDeepEqual(t *testing.T) {
	t.Run("when strings are deeply equal", shouldPass(func(t test.TestingT) {
		test.DeepEqual(t, "hello", "hello")
	}))

	t.Run("when strings are not deeply equal", shouldFail(func(t test.TestingT) {
		test.DeepEqual(t, "hello", "world")
	}))

	t.Run("when slices are deeply equal", shouldPass(func(t test.TestingT) {
		test.DeepEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	}))

	t.Run("when slices are not deeply equal", shouldFail(func(t test.TestingT) {
		test.DeepEqual(t, []int{1, 2, 3}, []int{1, 2, 4})
	}))

	t.Run("when maps are deeply equal", shouldPass(func(t test.TestingT) {
		test.DeepEqual(t,
			map[string]int{"a": 1, "b": 2},
			map[string]int{"a": 1, "b": 2},
		)
	}))

	t.Run("when maps have different values", shouldFail(func(t test.TestingT) {
		test.DeepEqual(t,
			map[string]int{"a": 1, "b": 2},
			map[string]int{"a": 1, "b": 3},
		)
	}))

	type person struct {
		Name string
		Age  int
	}

	t.Run("when structs are deeply equal", shouldPass(func(t test.TestingT) {
		test.DeepEqual(t,
			person{Name: "Alice", Age: 30},
			person{Name: "Alice", Age: 30},
		)
	}))

	t.Run("when structs have different values", shouldFail(func(t test.TestingT) {
		test.DeepEqual(t,
			person{Name: "Alice", Age: 30},
			person{Name: "Alice", Age: 31},
		)
	}))
}

func TestNotDeepEqual(t *testing.T) {
	t.Run("when slices are not deeply equal", shouldPass(func(t test.TestingT) {
		test.NotDeepEqual(t, []int{1, 2, 3}, []int{1, 2, 4})
	}))

	t.Run("when slices are deeply equal", shouldFail(func(t test.TestingT) {
		test.NotDeepEqual(t, []int{1, 2, 3}, []int{1, 2, 3})
	}))

	t.Run("when maps are not deeply equal", shouldPass(func(t test.TestingT) {
		test.NotDeepEqual(t,
			map[string]int{"a": 1, "b": 2},
			map[string]int{"a": 1, "b": 3},
		)
	}))

	type person struct {
		Name string
		Age  int
	}

	t.Run("when structs are not deeply equal", shouldPass(func(t test.TestingT) {
		test.NotDeepEqual(t,
			person{Name: "Alice", Age: 30},
			person{Name: "Alice", Age: 31},
		)
	}))
}
