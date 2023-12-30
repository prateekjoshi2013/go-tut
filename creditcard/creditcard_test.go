package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "1234567890"
	cc, err := creditcard.New(want)
	if err != nil {
		t.Errorf("New(%q) = %v, want nil", want, err)
	}
	got := cc.Number()
	if want != got {
		t.Errorf("New(%q).Number() = %q, want %q", want, got, want)
	}
}

func TestNewInvalidReturnsError(t *testing.T) {
	t.Parallel()
	_, err := creditcard.New("")
	if err == nil {
		t.Errorf("New(%q) = nil, want error", "12345678901234")
	}
}
