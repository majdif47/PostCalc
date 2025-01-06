package main

import "math"


	// List of unary operators
var constants = map[string]float64 {
  "pi": math.Pi,
  "π": math.Pi,
  "e": math.E,
  "phi": math.Phi,
  "φ": math.Phi,
}

var unaryOp = map[string]bool{
		"sqr": true, "sqrt": true, "log": true, "lg": true, "ln": true,
		"sin": true, "cos": true, "tan": true, "arcsin": true, "arctan": true,
		"fact": true, "!": true, "not": true,
	}

	// Map of operations
	var operations = map[string]func(a, b float64) float64{
		"+": func(a, b float64) float64 { return a + b },
		"-": func(a, b float64) float64 { return a - b },
		"x": func(a, b float64) float64 { return a * b },
		"/": func(a, b float64) float64 {
			if b == 0 {
				panic("Error: Division by zero")
			}
			return a / b
		},
		"root": func(a, b float64) float64 {
			if b == 0 {
				panic("Error: 0th root")
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
		"sin":  func(a, _ float64) float64 { return roundToZero(math.Sin(a))},
		"cos":  func(a, _ float64) float64 { return roundToZero(math.Cos(a))},
		"tan":  func(a, _ float64) float64 { return roundToZero(math.Tan(a))},
		"arcsin": func(a, _ float64) float64 { return roundToZero(math.Asin(a))},
		"arctan": func(a, _ float64) float64 { return roundToZero(math.Atan(a))},
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

