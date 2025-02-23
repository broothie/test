package match

import "fmt"

const (
	True  BoolMatcher = true
	False BoolMatcher = false
)

type BoolMatcher bool

func (m BoolMatcher) Match(subject bool) (string, bool) {
	return fmt.Sprintf("be %t", bool(m)), subject
}
