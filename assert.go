package test

import (
	"fmt"
)

func assert(t TestingT, assertion bool, expectedMessage string) bool {
	t.Helper()

	if !assertion {
		t.Error(fmt.Sprintf("expected %s", expectedMessage))
	}

	return assertion
}
