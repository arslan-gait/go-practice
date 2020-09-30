package myunpack //nolint:golint,stylecheck

import (
	"testing"
)

type test struct {
	input    string
	expected string
	err      error
}

//TestUnpack test w/o backslash
func TestUnpack(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abcd",
			expected: "abcd",
		},
		{
			input:    "45",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "aaa10b",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			input:    "",
			expected: "",
		},
	} {
		result, err := UnpackString(tst.input)
		if result != tst.expected {
			t.Errorf("expected %s, found %s, error %v", tst.expected, result, err)
		}
	}
}

//TestUnpackWithEscape test w backslash
func TestUnpackWithEscape(t *testing.T) {
	for _, tst := range [...]test{
		{
			input:    `qwe\4\5`,
			expected: `qwe45`,
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
		{
			input:    `qwe\\\3`,
			expected: `qwe\3`,
		},
	} {
		result, err := UnpackString(tst.input)
		if result != tst.expected || err != nil {
			t.Errorf("expected %s, found %s, error %v", tst.expected, result, err)
		}
	}
}
