package test

// Matcher is an interface for custom matchers that can match against a value.
type Matcher[Actual any] interface {
	Match(actual Actual) (bool, string)
}

// MatcherFunc is a function type that implements Matcher.
type MatcherFunc[Actual any] func(actual Actual) (bool, string)

func (f MatcherFunc[Actual]) Match(actual Actual) (bool, string) {
	return f(actual)
}

// Match asserts that a value matches a custom matcher.
func Match[Actual any](t TestingT, actual Actual, matcher Matcher[Actual]) bool {
	t.Helper()

	assertion, expectedMessage := matcher.Match(actual)
	return assert(t, assertion, expectedMessage)
}
