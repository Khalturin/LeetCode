package ReversString

import (
	"bytes"
	"testing"
)

func TestReversString1(t *testing.T) {
	s := "hello"
	expected := "olleh"
	res := ReversString([]byte(s))

	if !bytes.Equal([]byte(expected), res) {
		t.Error("fail TestReversString1, expected: ", expected, " result: ", string(res))
	}
}
func TestReversString2(t *testing.T) {
	s := "hannah"
	expected := "hannah"
	res := ReversString([]byte(s))

	if !bytes.Equal([]byte(expected), res) {
		t.Error("fail TestReversString2, expected: ", expected, " result: ", string(res))
	}
}
func TestReversString3(t *testing.T) {
	s := "abcdefg"
	expected := "gfedcba"
	res := ReversString([]byte(s))

	if !bytes.Equal([]byte(expected), res) {
		t.Error("fail TestReversString3, expected: ", expected, " result: ", string(res))
	}
}
