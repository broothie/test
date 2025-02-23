package match

import (
	"fmt"
	"regexp"
)

func Regexp(expected *regexp.Regexp) MatcherFunc[string] {
	return func(subject string) (string, bool) {
		return fmt.Sprintf("match %v", expected), expected.MatchString(subject)
	}
}
