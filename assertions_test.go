package test

import (
	"errors"
	"testing"

	"github.com/broothie/test/mocks"
	"go.uber.org/mock/gomock"
)

func Test_assertions(t *testing.T) {
	unwrappedError := errors.New("unwrapped")
	wrappedError := errors.Join(unwrappedError, errors.New("wrapped"))
	otherError := errors.New("other")

	type TestCase struct {
		run              func(t TestingT) bool
		mockExpectations func(mockTestingT *mocks.MockTestingT)
		expectedResult   bool
	}

	testCases := map[string]TestCase{
		"Equal pass": {
			run:            func(t TestingT) bool { return Equal(t, 10, 10) },
			expectedResult: true,
		},
		"Equal fail": {
			run: func(t TestingT) bool { return Equal(t, 10, 11) },
			mockExpectations: func(mockTestingT *mocks.MockTestingT) {
				mockTestingT.EXPECT().Errorf(errorfFormat, 10, 11)
			},
			expectedResult: false,
		},
		"True pass": {
			run:            func(t TestingT) bool { return True(t, true) },
			expectedResult: true,
		},
		"True fail": {
			run: func(t TestingT) bool { return True(t, false) },
			mockExpectations: func(mockTestingT *mocks.MockTestingT) {
				mockTestingT.EXPECT().Errorf(errorfFormat, false, true)
			},
			expectedResult: false,
		},
		"False pass": {
			run:            func(t TestingT) bool { return False(t, false) },
			expectedResult: true,
		},
		"False fail": {
			run: func(t TestingT) bool { return False(t, true) },
			mockExpectations: func(mockTestingT *mocks.MockTestingT) {
				mockTestingT.EXPECT().Errorf(errorfFormat, true, false)
			},
			expectedResult: false,
		},
		"ErrorIs pass": {
			run:            func(t TestingT) bool { return ErrorIs(t, wrappedError, unwrappedError) },
			expectedResult: true,
		},
		"ErrorIs fail": {
			run: func(t TestingT) bool { return ErrorIs(t, unwrappedError, otherError) },
			mockExpectations: func(mockTestingT *mocks.MockTestingT) {
				mockTestingT.EXPECT().Errorf(errorfFormat, unwrappedError, otherError)
			},
			expectedResult: false,
		},
		"NoError pass": {
			run:            func(t TestingT) bool { return NoError(t, nil) },
			expectedResult: true,
		},
		"NoError fail": {
			run: func(t TestingT) bool { return NoError(t, wrappedError) },
			mockExpectations: func(mockTestingT *mocks.MockTestingT) {
				mockTestingT.EXPECT().Errorf(errorfFormat, wrappedError, nil)
			},
			expectedResult: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testMock := newMockTestingT(t)

			if testCase.mockExpectations != nil {
				testCase.mockExpectations(testMock)
			}

			got := testCase.run(testMock)
			want := testCase.expectedResult
			if got != want {
				t.Errorf(errorfFormat, got, want)
			}
		})
	}
}

func TestMustNoError(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		testMock := newMockTestingT(t)

		MustNoError(testMock, nil)
	})

	t.Run("fail", func(t *testing.T) {
		err := errors.New("some error")

		testMock := newMockTestingT(t)
		testMock.EXPECT().Errorf(errorfFormat, err, nil)
		testMock.EXPECT().FailNow()

		MustNoError(testMock, err)
	})
}

func TestMust(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		testMock := newMockTestingT(t)

		want := "value"
		got := Must(testMock, func() (string, error) { return want, nil })
		if got != want {
			t.Errorf(errorfFormat, got, want)
		}
	})

	t.Run("fail", func(t *testing.T) {
		err := errors.New("some error")

		testMock := newMockTestingT(t)
		testMock.EXPECT().Errorf(errorfFormat, err, nil)
		testMock.EXPECT().FailNow()

		Must(testMock, func() (string, error) { return "value", err })
	})
}

func newMockTestingT(t *testing.T) *mocks.MockTestingT {
	testMock := mocks.NewMockTestingT(gomock.NewController(t))
	testMock.EXPECT().Helper().AnyTimes()

	return testMock
}
