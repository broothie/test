package test_test

import (
	"regexp"
	"testing"

	"github.com/broothie/test"
)

func TestRegexpMatch(t *testing.T) {
	t.Run("when string matches pattern exactly", shouldPass(func(t test.TestingT) {
		test.RegexpMatch(t, "hello", regexp.MustCompile(`^hello$`))
	}))

	t.Run("when string contains pattern", shouldPass(func(t test.TestingT) {
		test.RegexpMatch(t, "hello world", regexp.MustCompile(`world`))
	}))

	t.Run("when string matches pattern with wildcards", shouldPass(func(t test.TestingT) {
		test.RegexpMatch(t, "hello123", regexp.MustCompile(`hello\d+`))
	}))

	t.Run("when string does not match pattern", shouldFail(func(t test.TestingT) {
		test.RegexpMatch(t, "hello", regexp.MustCompile(`world`))
	}))

	t.Run("when string partially matches pattern", shouldFail(func(t test.TestingT) {
		test.RegexpMatch(t, "hello", regexp.MustCompile(`^world$`))
	}))
}

func TestNotRegexpMatch(t *testing.T) {
	t.Run("when string does not match pattern", shouldPass(func(t test.TestingT) {
		test.NotRegexpMatch(t, "hello", regexp.MustCompile(`world`))
	}))

	t.Run("when string matches pattern exactly", shouldFail(func(t test.TestingT) {
		test.NotRegexpMatch(t, "hello", regexp.MustCompile(`^hello$`))
	}))

	t.Run("when string contains pattern", shouldFail(func(t test.TestingT) {
		test.NotRegexpMatch(t, "hello world", regexp.MustCompile(`world`))
	}))

	t.Run("when string matches pattern with wildcards", shouldFail(func(t test.TestingT) {
		test.NotRegexpMatch(t, "hello123", regexp.MustCompile(`hello\d+`))
	}))
}
