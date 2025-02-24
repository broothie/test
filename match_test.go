package test_test

import (
	"testing"

	"github.com/broothie/test"
)

func TestMatch(t *testing.T) {
	type Person struct {
		ID   int
		Name string
	}

	samePerson := func(expected Person) test.MatcherFunc[Person] {
		return func(actual Person) (bool, string) {
			return actual.ID == expected.ID, "Persons to have the same ID"
		}
	}

	alice := Person{ID: 1, Name: "Alice"}
	bob := Person{ID: 2, Name: "Bob"}

	t.Run("when same person", shouldPass(func(t test.TestingT) {
		test.Match(t, alice, samePerson(alice))
	}))

	t.Run("when different person", shouldFail(func(t test.TestingT) {
		test.Match(t, alice, samePerson(bob))
	}))
}
