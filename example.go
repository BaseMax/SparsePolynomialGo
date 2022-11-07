package main

import (
	"fmt"
	"math/big"
)

// Main function
func main() {
	println("Hello, World!")

	// Create a sparse polynomial
	p := NewSparsePolynomial()
	p.Add(0, big.NewInt(1))
	p.Add(1, big.NewInt(2))
	p.Add(2, big.NewInt(3))
	p.Add(3, big.NewInt(4))
	p.Add(4, big.NewInt(5))
	p.Add(5, big.NewInt(6))

	// Print the sparse polynomial
	fmt.Println(p)

	// Evaluate the sparse polynomial at x = 2
	fmt.Println(p.Evaluate(big.NewInt(2)))

	// Create a sparse polynomial
	p2 := NewSparsePolynomial()
	p2.Add(0, big.NewInt(1))
	p2.Add(1, big.NewInt(2))
	p2.Add(2, big.NewInt(3))
	p2.Add(3, big.NewInt(4))
	p2.Add(4, big.NewInt(5))
	p2.Add(5, big.NewInt(6))

	// Print the sparse polynomial
	fmt.Println(p2)

	// Evaluate the sparse polynomial at x = 2
	fmt.Println(p2.Evaluate(big.NewInt(2)))

	// Addition
	fmt.Println(p.Addition(p2))

	// Subtraction
	fmt.Println(p.Subtraction(p2))

	// Multiplication
	fmt.Println(p.Multiplication(p2))

	// Division
	fmt.Println(p.Division(p2))

	// Modulus
	fmt.Println(p.Modulus(p2))

	// Integral
	fmt.Println(p.Integral())

	// Derivative
	fmt.Println(p.Derivative())

	// Inverse
	fmt.Println(p.Inverse())
}
