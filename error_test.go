package test_test

import (
	"errors"
	"fmt"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/broothie/test"
)

var (
	errBase    = errors.New("base error")
	errWrapped = fmt.Errorf("wrapped: %w", errBase)
)

func TestNoError(t *testing.T) {
	t.Run("when error is nil", shouldPass(func(t test.TestingT) {
		test.NoError(t, nil)
	}))

	t.Run("when error is not nil", shouldFail(func(t test.TestingT) {
		test.NoError(t, errBase)
	}))
}

func TestErrorMessageIs(t *testing.T) {
	t.Run("when error message matches exactly", shouldPass(func(t test.TestingT) {
		test.ErrorMessageIs(t, errors.New("foo"), "foo")
	}))

	t.Run("when error message does not match", shouldFail(func(t test.TestingT) {
		test.ErrorMessageIs(t, errors.New("foo"), "bar")
	}))
}

func TestNotErrorMessageIs(t *testing.T) {
	t.Run("when error message does not match", shouldPass(func(t test.TestingT) {
		test.NotErrorMessageIs(t, errors.New("foo"), "bar")
	}))

	t.Run("when error message matches exactly", shouldFail(func(t test.TestingT) {
		test.NotErrorMessageIs(t, errors.New("foo"), "foo")
	}))
}

func TestErrorIs(t *testing.T) {
	t.Run("when error wraps target error", shouldPass(func(t test.TestingT) {
		test.ErrorIs(t, errWrapped, errBase)
	}))

	t.Run("when error is target error", shouldPass(func(t test.TestingT) {
		test.ErrorIs(t, errBase, errBase)
	}))

	t.Run("when error does not wrap target error", shouldFail(func(t test.TestingT) {
		test.ErrorIs(t, errBase, errors.New("other"))
	}))
}

func TestNotErrorIs(t *testing.T) {
	t.Run("when error does not wrap target error", shouldPass(func(t test.TestingT) {
		test.NotErrorIs(t, errBase, errors.New("other"))
	}))

	t.Run("when error wraps target error", shouldFail(func(t test.TestingT) {
		test.NotErrorIs(t, errWrapped, errBase)
	}))

	t.Run("when error is target error", shouldFail(func(t test.TestingT) {
		test.NotErrorIs(t, errBase, errBase)
	}))
}

func TestMustNoError(t *testing.T) {
	t.Run("when error is nil", shouldPass(func(t test.TestingT) {
		test.MustNoError(t, nil)
	}))

	t.Run("when error is not nil", func(t *testing.T) {
		mockT := newMockT(t)
		mockT.EXPECT().Helper().AnyTimes()
		mockT.EXPECT().Error(gomock.AssignableToTypeOf(""))
		mockT.EXPECT().FailNow()

		test.MustNoError(mockT, errBase)
	})
}

func TestMust(t *testing.T) {
	t.Run("when function succeeds", shouldPass(func(t test.TestingT) {
		result := test.Must(t, func() (string, error) {
			return "success", nil
		})

		test.Equal(t, result, "success")
	}))

	t.Run("when function fails", func(t *testing.T) {
		mockT := newMockT(t)
		mockT.EXPECT().Helper().AnyTimes()
		mockT.EXPECT().Error(gomock.AssignableToTypeOf(""))
		mockT.EXPECT().FailNow()

		test.Must(mockT, func() (string, error) {
			return "", errBase
		})
	})
}
