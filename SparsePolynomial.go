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
			result.Add(d-d2, c.Div(c, c2))
		}
	}
	return result
}

// % operator for modulus two polynomial
func (p *SparsePolynomial) Modulus(p2 *SparsePolynomial) *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		for d2, c2 := range p2.coefficients {
			result.Add(d%d2, c.Mod(c, c2))
		}
	}
	return result
}

// Calculate integral of polynomial
func (p *SparsePolynomial) Integral() *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		result.Add(d+1, c.Div(c, big.NewInt(int64(d+1))))
	}
	return result
}

// Calculate derivative of polynomial
func (p *SparsePolynomial) Derivative() *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		result.Add(d-1, c.Mul(c, big.NewInt(int64(d))))
	}
	return result
}

// Calculate Inverse of polynomial
func (p *SparsePolynomial) Inverse() *SparsePolynomial {
	result := NewSparsePolynomial()
	for d, c := range p.coefficients {
		result.Add(d, c.Mul(c, big.NewInt(-1)))
	}
	return result
}
