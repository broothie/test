package match

import "fmt"

func Not[T any](matcher Matcher[T]) MatcherFunc[T] {
	return func(subject T) (string, bool) {
		to, isMatch := matcher.Match(subject)
		return fmt.Sprintf("not %s", to), !isMatch
	}
}
