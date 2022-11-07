/*
 * @Name: Sparse Polynomial Go
 * @Author: Max Base
 * @Date: 2022-11-07
 * @Class: Data Structure, Dr. Mahsa Soheil Shamaee
 * @Repository: https://github.com/basemax/SparsePolynomialGo
 */
package main

import (
	"fmt"
	"math/big"
)

// Sparse Polynomial type
type SparsePolynomial struct {
	coefficients map[int]*big.Int
}

// NewSparsePolynomial creates a new sparse polynomial
func NewSparsePolynomial() *SparsePolynomial {
	return &SparsePolynomial{make(map[int]*big.Int)}
}

// Add adds a coefficient to the sparse polynomial
func (p *SparsePolynomial) Add(degree int, coefficient *big.Int) {
	p.coefficients[degree] = coefficient
}

// Degree returns the degree of the sparse polynomial
func (p *SparsePolynomial) Degree() int {
	degree := 0
	for d := range p.coefficients {
		if d > degree {
			degree = d
		}
	}
	return degree
}

// Evaluate evaluates the sparse polynomial at a given point
func (p *SparsePolynomial) Evaluate(x *big.Int) *big.Int {
	result := big.NewInt(0)
	for d, c := range p.coefficients {
		result.Add(result, c.Mul(c, new(big.Int).Exp(x, big.NewInt(int64(d)), nil)))
	}
	return result
}

// String returns a string representation of the sparse polynomial
func (p *SparsePolynomial) String() string {
	result := ""
	for d, c := range p.coefficients {
		result += fmt.Sprintf("%d*x^%d + ", c, d)
	}
	return result[:len(result)-3]
}

// + operator for sum two polynomial
func (p *SparsePolynomial) Addition(p2 *SparsePolynomial) *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		result.Add(d, c)
	}
	for d, c := range p2.coefficients {
		result.Add(d, c)
	}
	return result
}

// - operator for subtract two polynomial
func (p *SparsePolynomial) Subtraction(p2 *SparsePolynomial) *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		result.Add(d, c)
	}
	for d, c := range p2.coefficients {
		result.Add(d, c.Neg(c))
	}
	return result
}

// * operator for multiply two polynomial
func (p *SparsePolynomial) Multiplication(p2 *SparsePolynomial) *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		for d2, c2 := range p2.coefficients {
			result.Add(d+d2, c.Mul(c, c2))
		}
	}
	return result
}

// / operator for divide two polynomial
func (p *SparsePolynomial) Division(p2 *SparsePolynomial) *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		for d2, c2 := range p2.coefficients {
			result.Add(d+d2, c.Mul(c, c2))
		}
	}
	return result
}

// % operator for modulus two polynomial
func (p *SparsePolynomial) Modulus(p2 *SparsePolynomial) *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		for d2, c2 := range p2.coefficients {
			result.Add(d+d2, c.Mul(c, c2))
		}
	}
	return result
}

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
}
