// Package adder provides basic mathematical operations.
// This package includes a simple function to add two integers.

package adder

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two integers and returns their sum.
// This function performs a basic addition operation.
// For more information about addition, you can visit:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
