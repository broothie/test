package test

//go:generate go run go.uber.org/mock/mockgen@v0.5.0 --destination mocktesting/testing_t.go --package mocktesting . TestingT
type TestingT interface {
	Helper()
	Error(...any)
	FailNow()
}
