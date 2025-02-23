package match

import (
	"cmp"
	"fmt"
)

func LessThan[T cmp.Ordered](expected T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("be < %v", expected), subject < expected
	}
}

func LessThanOrEqual[T cmp.Ordered](expected T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("be <= %s", expected), subject <= expected
	}
}

func GreaterThan[T cmp.Ordered](expected T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("be > %v", expected), subject > expected
	}
}

func GreaterThanOrEqual[T cmp.Ordered](expected T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("be >= %v", expected), subject >= expected
	}
}

func BetweenInclusive[T cmp.Ordered](minimum, maximum T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("be >= %v and <= %v", minimum, maximum), minimum <= subject && subject <= maximum
	}
}

func BetweenExclusive[T cmp.Ordered](minimum, maximum T) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		return fmt.Sprintf("be > %v and < %v", minimum, maximum), minimum < subject && subject < maximum
	}
}
