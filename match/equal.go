package match

import (
	"fmt"
	"reflect"
)

func Equal[T comparable](expected T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("== %v", expected), subject == expected
	}
}

func DeepEqual[T any](expected T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("deep equal %v", expected), reflect.DeepEqual(subject, expected)
	}
}
