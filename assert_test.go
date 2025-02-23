package test

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/broothie/test/mocktesting"
)

func TestAssert(t *testing.T) {
	expectedMessage := "expected message"

	mockT := mocktesting.NewMockTestingT(gomock.NewController(t))
	mockT.EXPECT().Helper().AnyTimes()
	mockT.EXPECT().
		Error(gomock.AssignableToTypeOf("")).
		DoAndReturn(func(message string) {
			if message != expectedMessage {
				t.Errorf("got: %v, want: %v", message, expectedMessage)
			}
		})

	assert(mockT, false, "message")
}
