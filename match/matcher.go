package match

type Matcher[Subject any] interface {
	Match(subject Subject) (to string, isMatch bool)
}

type MatcherFunc[Subject any] func(subject Subject) (string, bool)

func (f MatcherFunc[Subject]) Match(subject Subject) (string, bool) {
	return f(subject)
}
