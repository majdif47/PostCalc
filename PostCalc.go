package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	stack "github.com/majdif47/Go-dsa/ds/stacks"
)

// Factorial function
func fact(num int) int {
	if num <= 1 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	// List of unary operators
	unaryOp := map[string]bool{
		"sqr": true, "sqrt": true, "log": true, "lg": true, "ln": true,
		"sin": true, "cos": true, "tan": true, "arcsin": true, "arctan": true,
		"fact": true, "!": true, "not": true,
	}

	// Map of operations
	operations := map[string]func(a, b float64) float64{
		"+": func(a, b float64) float64 { return a + b },
		"-": func(a, b float64) float64 { return a - b },
		"x": func(a, b float64) float64 { return a * b },
		"/": func(a, b float64) float64 {
			if b == 0 {
				fmt.Println("Error: Division by zero")
				return math.NaN()
			}
			return a / b
		},
		"root": func(a, b float64) float64 {
			if b == 0 {
				fmt.Println("Error: 0th root")
				return math.NaN()
			}
			return math.Pow(a, 1/b)
		},
		"pow":  func(a, b float64) float64 { return math.Pow(a, b) },
		"**":   func(a, b float64) float64 { return math.Pow(a, b) },
		"sqr":  func(a, _ float64) float64 { return math.Pow(a, 2) },
		"sqrt": func(a, _ float64) float64 { return math.Sqrt(a) },
		"log":  func(a, _ float64) float64 { return math.Log10(a) },
		"lg":   func(a, _ float64) float64 { return math.Log2(a) },
		"ln":   func(a, _ float64) float64 { return math.Log(a) },
		"sin":  func(a, _ float64) float64 { return math.Sin(a) },
		"cos":  func(a, _ float64) float64 { return math.Cos(a) },
		"tan":  func(a, _ float64) float64 { return math.Tan(a) },
		"arcsin": func(a, _ float64) float64 { return math.Asin(a) },
		"arctan": func(a, _ float64) float64 { return math.Atan(a) },
		"%": func(a, b float64) float64 {
			return float64(int(a) % int(b))
		},
		"fact": func(a, _ float64) float64 {
			return float64(fact(int(a)))
		},
		"!": func(a, _ float64) float64 {
			return float64(fact(int(a)))
		},
		"and": func(a, b float64) float64 {
			return float64(int(a) & int(b))
		},
		"&": func(a, b float64) float64 {
			return float64(int(a) & int(b))
		},
		"or": func(a, b float64) float64 {
			return float64(int(a) | int(b))
		},
		"|": func(a, b float64) float64 {
			return float64(int(a) | int(b))
		},
		"xor": func(a, b float64) float64 {
			return float64(int(a) ^ int(b))
		},
		"^": func(a, b float64) float64 {
			return float64(int(a) ^ int(b))
		},
		"not": func(a, _ float64) float64 {
			return float64(^int(a))
		},
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: pcalc <expression>")
		return
	}

	args := os.Args[1:]
	s := stack.NewStack[float64]()

	for _, arg := range args {
		if num, err := strconv.ParseFloat(arg, 64); err == nil {
			// Push numbers to stack
			s.Push(num)
		} else {
			// Check for operator
			if opFunc, exists := operations[arg]; exists {
				if unaryOp[arg] {
					// Unary operator
					a, _ := s.Pop()
					result := opFunc(a, 0)
					s.Push(result)
				} else {
					// Binary operator
					b, _ := s.Pop()
					a, _ := s.Pop()
					result := opFunc(a, b)
					s.Push(result)
				}
			} else {
				fmt.Println("Error: Unknown operator", arg)
				return
			}
		}
	}

	// Final result
	if s.Size() != 1 {
		fmt.Println("Error: Invalid expression")
		return
	}
	result, _ := s.Pop()
	fmt.Println(result)
}

