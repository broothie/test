package test

//go:generate go run go.uber.org/mock/mockgen@v0.5.0 --destination mocktesting/testing_t.go --package mocktesting . TestingT
// TestingT is an interface wrapper around *testing.T.
type TestingT interface {
	Helper()
	Error(args ...any)
	FailNow()
}
