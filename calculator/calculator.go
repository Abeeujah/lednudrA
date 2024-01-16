// Package calculator does simple calculations
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding
// Two of them
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers a and b and
// returns the result of subtracting b from a
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers a and b and
// returns the result of multiplying a and b
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers a and b and
// returns the result of dividing a and b
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero not allowed")
	}
	return a / b, nil
}

// Sqrt takes a number a and
// returns the square root of a
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Square root of negative numbers not allowed")
	}
	return math.Sqrt(a), nil
}
