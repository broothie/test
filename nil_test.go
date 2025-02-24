package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestNil(t *testing.T) {
	t.Run("when value is nil", shouldPass(func(t test.TestingT) {
		test.Nil(t, nil)
	}))

	t.Run("when value is not nil", shouldFail(func(t test.TestingT) {
		test.Nil(t, "not nil")
	}))

	t.Run("when slice is nil", shouldPass(func(t test.TestingT) {
		var slice []int = nil
		test.Nil(t, slice)
	}))

	t.Run("when map is nil", shouldPass(func(t test.TestingT) {
		var m map[string]int = nil
		test.Nil(t, m)
	}))
}

func TestNotNil(t *testing.T) {
	t.Run("when value is not nil", shouldPass(func(t test.TestingT) {
		test.NotNil(t, "not nil")
	}))

	t.Run("when value is nil", shouldFail(func(t test.TestingT) {
		test.NotNil(t, nil)
	}))

	t.Run("when slice is not nil", shouldPass(func(t test.TestingT) {
		test.NotNil(t, []int{})
	}))

	t.Run("when map is not nil", shouldPass(func(t test.TestingT) {
		test.NotNil(t, map[string]int{})
	}))
}
