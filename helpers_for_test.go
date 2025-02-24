package test_test

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/broothie/test"
	"github.com/broothie/test/mocktesting"
)

func newMockT(t *testing.T) *mocktesting.MockTestingT {
	return mocktesting.NewMockTestingT(gomock.NewController(t))
}

func shouldPass(f func(t test.TestingT)) func(t *testing.T) {
	return func(t *testing.T) {
		mockT := newMockT(t)
		mockT.EXPECT().Helper().AnyTimes()

		f(mockT)
	}
}

func shouldFail(f func(t test.TestingT)) func(t *testing.T) {
	return func(t *testing.T) {
		mockT := newMockT(t)
		mockT.EXPECT().Helper().AnyTimes()
		mockT.EXPECT().Error(gomock.AssignableToTypeOf(""))
		f(mockT)
	}
}
