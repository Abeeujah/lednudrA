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

// AddMany takes multiple numbers and returns their sum
func AddMany(nums ...float64) float64 {
	sum := float64(0)
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Subtract takes two numbers a and b and
// returns the result of subtracting b from a
func Subtract(a, b float64) float64 {
	return a - b
}

// SubtractMany takes multiple numbers and returns their
// difference
func SubtractMany(nums ...float64) float64 {
	diff := nums[0]
	for index, num := range nums {
		if index == 0 {
			continue
		}
		diff -= num
	}
	return diff
}

// Multiply takes two numbers a and b and
// returns the result of multiplying a and b
func Multiply(a, b float64) float64 {
	return a * b
}

// MultiplyMany takes many numbers and returns
// their product
func MultiplyMany(nums ...float64) float64 {
	product := float64(1)
	for _, num := range nums {
		product *= num
	}
	return product
}

// Divide takes two numbers a and b and
// returns the result of dividing a and b
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Division by zero not allowed.")
	}
	return a / b, nil
}

// DivideMany takes multiple numbers and returns
// the result of dividing them
func DivideMany(nums ...float64) (float64, error) {
	result := nums[0]
	for index, num := range nums {
		if index == 0 {
			continue
		}
		if num == float64(0) {
			return 0, errors.New("Division by zero not allowed.")
		}
		result /= num
	}
	return result, nil
}

// Sqrt takes a number a and
// returns the square root of a
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Square root of negative numbers not allowed")
	}
	return math.Sqrt(a), nil
}
