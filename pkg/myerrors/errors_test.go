package myerrors

import (
	"errors"
	"fmt"
	"slices"
	"testing"
)

var (
	testErr1 = errors.New("1")
	testErr2 = errors.New("2")
	testErr3 = errors.New("3")
	testErr4 = errors.New("4")
)

func TestJoin(t *testing.T) {
	tests := []struct {
		input   []error
		wantMsg string
	}{
		{nil, ""},
		{[]error{}, ""},
		{[]error{testErr1}, testErr1.Error()},
		{[]error{testErr1, testErr2}, fmt.Sprintf("%v%s%v", testErr1, delimiter, testErr2)},
		{[]error{testErr1, testErr1}, fmt.Sprintf("%v%s%v", testErr1, delimiter, testErr1)},
		{[]error{testErr1, testErr2, testErr3, testErr4}, fmt.Sprintf("%v%s%v%s%v%s%v", testErr1, delimiter, testErr2, delimiter, testErr3, delimiter, testErr4)},
		{[]error{testErr1, nil, testErr2}, fmt.Sprintf("%v%s%s%v", testErr1, delimiter, delimiter, testErr2)},
	}

	for _, test := range tests {
		gotErr := Join(test.input...)
		errWrap, ok := (gotErr).(*errorWrapper)
		if !ok {
			t.Errorf("Join(%v) = %v. Wrong type", test.input, gotErr)
			continue
		}
		if !slices.Equal(test.input, errWrap.wrappedErrors) || errWrap.msg != test.wantMsg {
			t.Errorf("Join(%v) = %v.\nWant: %v", test.input, errWrap, &errorWrapper{msg: test.wantMsg, wrappedErrors: test.input})
		}
	}
}

func TestExtractWrapped(t *testing.T) {
	tests := []struct {
		input error
		want  []string
	}{
		{nil, nil},
		{testErr1, []string{testErr1.Error()}},
		{fmt.Errorf("error: %w", testErr1), []string{fmt.Errorf("error: %w", testErr1).Error()}},
		{Join(testErr1), []string{testErr1.Error()}},
		{Join(testErr1, testErr2, testErr3), []string{testErr1.Error(), testErr2.Error(), testErr3.Error()}},
		{Join(testErr1, nil, testErr2), []string{testErr1.Error(), testErr2.Error()}},
		{fmt.Errorf("%w %w %w", testErr3, nil, testErr4), []string{testErr3.Error(), testErr4.Error()}},
	}

	for _, test := range tests {
		if got := ExtractWrapped(test.input); !slices.Equal(got, test.want) {
			t.Errorf("ExtractWrapped(%v) = %v.\nWant: %v", test.input, got, test.want)
		}
	}
}
