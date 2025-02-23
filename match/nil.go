package match

import "fmt"

const Nil NilMatcher = 0

type NilMatcher int

func (m NilMatcher) Match(subject any) (string, bool) {
	return fmt.Sprintf("be nil"), subject == nil
}
