package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestLessThan(t *testing.T) {
	t.Run("when value is less than expected", shouldPass(func(t test.TestingT) {
		test.LessThan(t, 1, 2)
	}))

	t.Run("when value is equal to expected", shouldFail(func(t test.TestingT) {
		test.LessThan(t, 2, 2)
	}))

	t.Run("when value is greater than expected", shouldFail(func(t test.TestingT) {
		test.LessThan(t, 3, 2)
	}))
}

func TestLessThanOrEqual(t *testing.T) {
	t.Run("when value is less than expected", shouldPass(func(t test.TestingT) {
		test.LessThanOrEqual(t, 1, 2)
	}))

	t.Run("when value is equal to expected", shouldPass(func(t test.TestingT) {
		test.LessThanOrEqual(t, 2, 2)
	}))

	t.Run("when value is greater than expected", shouldFail(func(t test.TestingT) {
		test.LessThanOrEqual(t, 3, 2)
	}))
}

func TestGreaterThan(t *testing.T) {
	t.Run("when value is greater than expected", shouldPass(func(t test.TestingT) {
		test.GreaterThan(t, 3, 2)
	}))

	t.Run("when value is equal to expected", shouldFail(func(t test.TestingT) {
		test.GreaterThan(t, 2, 2)
	}))

	t.Run("when value is less than expected", shouldFail(func(t test.TestingT) {
		test.GreaterThan(t, 1, 2)
	}))
}

func TestGreaterThanOrEqual(t *testing.T) {
	t.Run("when value is greater than expected", shouldPass(func(t test.TestingT) {
		test.GreaterThanOrEqual(t, 3, 2)
	}))

	t.Run("when value is equal to expected", shouldPass(func(t test.TestingT) {
		test.GreaterThanOrEqual(t, 2, 2)
	}))

	t.Run("when value is less than expected", shouldFail(func(t test.TestingT) {
		test.GreaterThanOrEqual(t, 1, 2)
	}))
}

func TestBetweenInclusive(t *testing.T) {
	t.Run("when value is equal to lower bound", shouldPass(func(t test.TestingT) {
		test.BetweenInclusive(t, 1, 1, 3)
	}))

	t.Run("when value is between bounds", shouldPass(func(t test.TestingT) {
		test.BetweenInclusive(t, 2, 1, 3)
	}))

	t.Run("when value is equal to upper bound", shouldPass(func(t test.TestingT) {
		test.BetweenInclusive(t, 3, 1, 3)
	}))

	t.Run("when value is less than lower bound", shouldFail(func(t test.TestingT) {
		test.BetweenInclusive(t, 0, 1, 3)
	}))

	t.Run("when value is greater than upper bound", shouldFail(func(t test.TestingT) {
		test.BetweenInclusive(t, 4, 1, 3)
	}))
}

func TestBetweenExclusive(t *testing.T) {
	t.Run("when value is equal to lower bound", shouldFail(func(t test.TestingT) {
		test.BetweenExclusive(t, 1, 1, 3)
	}))

	t.Run("when value is between bounds", shouldPass(func(t test.TestingT) {
		test.BetweenExclusive(t, 2, 1, 3)
	}))

	t.Run("when value is equal to upper bound", shouldFail(func(t test.TestingT) {
		test.BetweenExclusive(t, 3, 1, 3)
	}))

	t.Run("when value is less than lower bound", shouldFail(func(t test.TestingT) {
		test.BetweenExclusive(t, 0, 1, 3)
	}))

	t.Run("when value is greater than upper bound", shouldFail(func(t test.TestingT) {
		test.BetweenExclusive(t, 4, 1, 3)
	}))
}
