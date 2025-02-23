package match

import "fmt"

func PanicWith(expected any) MatcherFunc[func()] {
	return func(subject func()) (string, bool) {
		got := captureRecovery(subject)

		return fmt.Sprintf("panic with %v (got %v)", expected, got), got == expected
	}
}

func captureRecovery(f func()) (value any) {
	defer func() { value = recover() }()

	f()
	return
}
