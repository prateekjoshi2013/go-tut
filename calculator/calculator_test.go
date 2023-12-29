package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testcase struct {
	a, b float64
	want float64
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testcases := []testcase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testcases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {

			t.Errorf("want %f got %f", tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testcases := []testcase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 1, want: 0},
		{a: 5, b: 0, want: -5},
	}
	for _, tc := range testcases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("got %f want %f", got, tc.want)
		}
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testcases := []testcase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
	}
	for _, tc := range testcases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("got %f want %f", got, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testcases := []testcase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.33333},
	}
	for _, tc := range testcases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(got, tc.want, 0.001) {
			t.Errorf("Divide(%f,%f), got: %f want %f", tc.a, tc.b, got, tc.want)
		}
	}

}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	got, err := calculator.Divide(2, 0)
	if err == nil {
		t.Errorf("expected error got %v", got)
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testcases := []testcase{
		{4, 0, 2},
		{9, 0, 3},
		{2, 0, 1.414},
	}
	for _, tc := range testcases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("error not expected")
		}
		if !closeEnough(got, tc.want, 0.001) {
			t.Fatalf("got %f want %f", got, tc.want)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	got, err := calculator.Sqrt(-2)
	if err == nil {
		t.Errorf("expected error got %v", got)
	}

}
