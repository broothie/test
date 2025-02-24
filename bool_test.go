package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestTrue(t *testing.T) {
	t.Run("when value is true", shouldPass(func(t test.TestingT) {
		test.True(t, true)
	}))

	t.Run("when value is false", shouldFail(func(t test.TestingT) {
		test.True(t, false)
	}))
}

func TestFalse(t *testing.T) {
	t.Run("when value is false", shouldPass(func(t test.TestingT) {
		test.False(t, false)
	}))

	t.Run("when value is true", shouldFail(func(t test.TestingT) {
		test.False(t, true)
	}))
}
