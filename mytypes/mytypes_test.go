package mytypes_test

import (
	"mytypes"
	"strings"
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

func TestStringsBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("world!")
	want := "Hello, world!"
	got := sb.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
	wantLen := len(want)
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("want %d, got %d", wantLen, gotLen)
	}
}

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}

}

func TestMyBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("world!")
	want := "Hello, world!"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestStringUpperCaser(t *testing.T) {
	t.Parallel()
	var sb mytypes.StringUpperCaser
	sb.Contents.WriteString("Hello, ")
	sb.Contents.WriteString("world!")
	want := "HELLO, WORLD!"
	got := sb.ToUpper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	x := mytypes.MyInt(12)
	want := mytypes.MyInt(24)
	p := &x
	p.Double()
	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}

}
