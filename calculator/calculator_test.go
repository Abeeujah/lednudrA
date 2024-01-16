package calculator_test

import (
	"calculator"
	"errors"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): expected %f, got %f\n", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 8, b: 3, want: 5},
		{a: 9, b: 2, want: 7},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): expected %f, got %f\n", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 3, b: 3, want: 9},
		{a: 4, b: 4, want: 16},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): expected %f, got %f\n", tc.a, tc.b, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: -1, b: -1, want: 1},
		{a: 10, b: -2, want: -5},
		{a: 1, b: 3, want: 0.33333},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("Expected no error for valid input, got %v\n", err)
		}

		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): expected %f, got %f\n", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want error
	}
	testCases := []testCase{
		{a: 0, b: 0, want: errors.New("Division by zero not allowed")},
		{a: 5, b: 0, want: errors.New("Division by zero not allowed")},
		{a: 10, b: -0, want: errors.New("Division by zero not allowed")},
	}

	for _, tc := range testCases {
		_, err := calculator.Divide(tc.a, tc.b)

		if err == nil {
			t.Errorf("Divide(%f, %f): expected %f, got %f\n", tc.a, tc.b, tc.want, err)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 74, want: 8.602323},
		{a: 9, want: 3},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		if err != nil {
			t.Fatalf("Error value cannot be gotten for valid inputs, got %v\n", err)
		}

		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("Sqrt(%f): expected %f, got %f\n", tc.a, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    float64
		want error
	}
	testCases := []testCase{
		{a: -4, want: errors.New("Square root of negative numbers not allowed")},
		{a: -12, want: errors.New("Square root of negative numbers not allowed")},
		{a: -48, want: errors.New("Square root of negative numbers not allowed")},
	}

	for _, tc := range testCases {
		_, err := calculator.Sqrt(tc.a)

		if err == nil {
			t.Errorf("Sqrt(%f): expected %v, got %v\n", tc.a, tc.want, err)
		}
	}
}
