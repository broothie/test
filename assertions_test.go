package test_test

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/broothie/test"
	"github.com/broothie/test/mocktesting"
)

const errorfFormat = "got: %+v, want: %+v"

func Test_assertions(t *testing.T) {
	unwrappedError := errors.New("unwrapped")
	wrappedError := errors.Join(unwrappedError, errors.New("wrapped"))
	otherError := errors.New("other")

	type TestCase struct {
		run              func(t test.TestingT) bool
		mockExpectations func(mockTestingT *mocktesting.MockTestingT)
		expectedResult   bool
	}

	testCases := map[string]TestCase{
		"DeepEqual pass": {
			run:            func(t test.TestingT) bool { return test.DeepEqual(t, map[string]any{"a": 1}, map[string]any{"a": 1}) },
			expectedResult: true,
		},
		"DeepEqual false": {
			run: func(t test.TestingT) bool { return test.DeepEqual(t, map[string]any{"a": 1}, map[string]any{"a": 2}) },
			mockExpectations: func(mockTestingT *mocktesting.MockTestingT) {
				mockTestingT.EXPECT().Error(gomock.AssignableToTypeOf(""))
			},
			expectedResult: false,
		},
		"Equal pass": {
			run:            func(t test.TestingT) bool { return test.Equal(t, 10, 10) },
			expectedResult: true,
		},
		"Equal fail": {
			run: func(t test.TestingT) bool { return test.Equal(t, 10, 11) },
			mockExpectations: func(mockTestingT *mocktesting.MockTestingT) {
				mockTestingT.EXPECT().Error(gomock.AssignableToTypeOf(""))
			},
			expectedResult: false,
		},
		"True pass": {
			run:            func(t test.TestingT) bool { return test.True(t, true) },
			expectedResult: true,
		},
		"True fail": {
			run: func(t test.TestingT) bool { return test.True(t, false) },
			mockExpectations: func(mockTestingT *mocktesting.MockTestingT) {
				mockTestingT.EXPECT().Error(gomock.AssignableToTypeOf(""))
			},
			expectedResult: false,
		},
		"False pass": {
			run:            func(t test.TestingT) bool { return test.False(t, false) },
			expectedResult: true,
		},
		"False fail": {
			run: func(t test.TestingT) bool { return test.False(t, true) },
			mockExpectations: func(mockTestingT *mocktesting.MockTestingT) {
				mockTestingT.EXPECT().Error(gomock.AssignableToTypeOf(""))
			},
			expectedResult: false,
		},
		"ErrorIs pass": {
			run:            func(t test.TestingT) bool { return test.ErrorIs(t, wrappedError, unwrappedError) },
			expectedResult: true,
		},
		"ErrorIs fail": {
			run: func(t test.TestingT) bool { return test.ErrorIs(t, unwrappedError, otherError) },
			mockExpectations: func(mockTestingT *mocktesting.MockTestingT) {
				mockTestingT.EXPECT().Error(gomock.AssignableToTypeOf(""))
			},
			expectedResult: false,
		},
		"Nil pass": {
			run:            func(t test.TestingT) bool { return test.Nil(t, nil) },
			expectedResult: true,
		},
		"Nil fail": {
			run: func(t test.TestingT) bool { return test.Nil(t, wrappedError) },
			mockExpectations: func(mockTestingT *mocktesting.MockTestingT) {
				mockTestingT.EXPECT().Error(gomock.AssignableToTypeOf(""))
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

		test.MustNoError(testMock, nil)
	})

	t.Run("fail", func(t *testing.T) {
		err := errors.New("some error")

		testMock := newMockTestingT(t)
		testMock.EXPECT().Error(gomock.AssignableToTypeOf(""))
		testMock.EXPECT().FailNow()

		test.MustNoError(testMock, err)
	})
}

func TestMust(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		testMock := newMockTestingT(t)

		want := "value"
		got := test.Must(testMock, func() (string, error) { return want, nil })
		if got != want {
			t.Errorf(errorfFormat, got, want)
		}
	})

	t.Run("fail", func(t *testing.T) {
		err := errors.New("some error")

		testMock := newMockTestingT(t)
		testMock.EXPECT().Error(gomock.AssignableToTypeOf(""))
		testMock.EXPECT().FailNow()

		test.Must(testMock, func() (string, error) { return "value", err })
	})
}

func newMockTestingT(t *testing.T) *mocktesting.MockTestingT {
	testMock := mocktesting.NewMockTestingT(gomock.NewController(t))
	testMock.EXPECT().Helper().AnyTimes()

	return testMock
}
