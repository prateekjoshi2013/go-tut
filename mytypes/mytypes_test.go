package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9) // casting to MyInt
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("twice %d: want %d, got %d", input, want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	str := "mystring"
	input := mytypes.MyString(str)
	want := len(str)
	got := input.Len()
	if got != want {
		t.Errorf("length for %s got %d want %d", str, got, want)
	}
}
