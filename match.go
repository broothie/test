package test

type Matcher[Actual any] interface {
	Match(actual Actual) (bool, string)
}

type MatcherFunc[Actual any] func(actual Actual) (bool, string)

func (f MatcherFunc[Actual]) Match(actual Actual) (bool, string) {
	return f(actual)
}

func Match[Actual any](t TestingT, actual Actual, matcher Matcher[Actual]) bool {
	t.Helper()

	assertion, expectedMessage := matcher.Match(actual)
	return assert(t, assertion, expectedMessage)
}
