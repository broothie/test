package match

import (
	"errors"
	"fmt"
)

func IsError() MatcherFunc[any] {
	return func(subject any) (string, bool) {
		_, isMatch := subject.(error)
		return "be an error", isMatch
	}
}

func ErrorIs(expected error) MatcherFunc[error] {
	return func(subject error) (string, bool) {
		return fmt.Sprintf("match %v", expected), errors.Is(subject, expected)
	}
}

func ErrorMessage(expected string) MatcherFunc[error] {
	return func(subject error) (string, bool) {
		return fmt.Sprintf("have error message %q", expected), subject.Error() == expected
	}
}
