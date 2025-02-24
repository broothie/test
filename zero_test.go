package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestZero(t *testing.T) {
	t.Run("when int is zero", shouldPass(func(t test.TestingT) {
		test.Zero(t, 0)
	}))

	t.Run("when int is not zero", shouldFail(func(t test.TestingT) {
		test.Zero(t, 42)
	}))

	t.Run("when string is zero", shouldPass(func(t test.TestingT) {
		test.Zero(t, "")
	}))

	t.Run("when string is not zero", shouldFail(func(t test.TestingT) {
		test.Zero(t, "hello")
	}))

	t.Run("when bool is zero", shouldPass(func(t test.TestingT) {
		test.Zero(t, false)
	}))

	t.Run("when bool is not zero", shouldFail(func(t test.TestingT) {
		test.Zero(t, true)
	}))

	type Person struct {
		Name string
		Age  int
	}

	t.Run("when struct is zero", shouldPass(func(t test.TestingT) {
		test.Zero(t, Person{})
	}))

	t.Run("when struct is not zero", shouldFail(func(t test.TestingT) {
		test.Zero(t, Person{Name: "Alice", Age: 30})
	}))

	t.Run("when slice is zero", shouldPass(func(t test.TestingT) {
		test.Zero(t, []int(nil))
	}))

	t.Run("when slice is not zero", shouldFail(func(t test.TestingT) {
		test.Zero(t, []int{1, 2, 3})
	}))

	t.Run("when map is zero", shouldPass(func(t test.TestingT) {
		test.Zero(t, map[string]int(nil))
	}))

	t.Run("when map is not zero", shouldFail(func(t test.TestingT) {
		test.Zero(t, map[string]int{"foo": 1})
	}))

	t.Run("when pointer is not zero", shouldFail(func(t test.TestingT) {
		i := 1
		test.Zero(t, &i)
	}))

	t.Run("when pointer is zero", shouldPass(func(t test.TestingT) {
		test.Zero(t, (*int)(nil))
	}))
}

func TestNotZero(t *testing.T) {
	t.Run("when int is not zero", shouldPass(func(t test.TestingT) {
		test.NotZero(t, 42)
	}))

	t.Run("when int is zero", shouldFail(func(t test.TestingT) {
		test.NotZero(t, 0)
	}))

	t.Run("when string is not zero", shouldPass(func(t test.TestingT) {
		test.NotZero(t, "hello")
	}))

	t.Run("when string is zero", shouldFail(func(t test.TestingT) {
		test.NotZero(t, "")
	}))

	t.Run("when bool is not zero", shouldPass(func(t test.TestingT) {
		test.NotZero(t, true)
	}))

	t.Run("when bool is zero", shouldFail(func(t test.TestingT) {
		test.NotZero(t, false)
	}))

	type Person struct {
		Name string
		Age  int
	}

	t.Run("when struct is not zero", shouldPass(func(t test.TestingT) {
		test.NotZero(t, Person{Name: "Alice", Age: 30})
	}))

	t.Run("when struct is zero", shouldFail(func(t test.TestingT) {
		test.NotZero(t, Person{})
	}))

	t.Run("when slice is not zero", shouldPass(func(t test.TestingT) {
		test.NotZero(t, []int{1, 2, 3})
	}))

	t.Run("when slice is zero", shouldFail(func(t test.TestingT) {
		test.NotZero(t, []int(nil))
	}))

	t.Run("when map is not zero", shouldPass(func(t test.TestingT) {
		test.NotZero(t, map[string]int{"foo": 1})
	}))

	t.Run("when map is zero", shouldFail(func(t test.TestingT) {
		test.NotZero(t, map[string]int(nil))
	}))

	t.Run("when pointer is not zero", shouldPass(func(t test.TestingT) {
		i := 1
		test.NotZero(t, &i)
	}))

	t.Run("when pointer is zero", shouldFail(func(t test.TestingT) {
		test.NotZero(t, (*int)(nil))
	}))
}
