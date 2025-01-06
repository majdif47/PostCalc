# 🔢 Reverse Polish Notation (RPN) Calculator

A sleek and efficient command-line calculator built in Go to evaluate mathematical expressions using **Reverse Polish Notation (RPN)**. Supports a wide range of operations including arithmetic, trigonometry, and logical functions, all with beautifully styled output. ✨

---

## 🤔 What is Reverse Polish Notation?

**Reverse Polish Notation (RPN)** is a mathematical notation where operators follow their operands. Unlike standard infix notation (e.g., `3 + 5`), RPN does not require parentheses to dictate operation order.  
For example:
- **Infix**: `3 + 5`  
- **RPN**: `3 5 +`

RPN evaluates expressions efficiently using a stack, making it a popular choice for calculators and programming languages. 🧠

---

## 🌟 Features

- 🧮 **Arithmetic Operations**:  
  `+`, `-`, `x`, `/`, `%`
- 🔢 **Advanced Mathematical Operations**:  
  `pow` (Power), `root` (Root), `sqr` (Square), `sqrt` (Square Root), `log` (Logarithm base 10), `ln` (Natural Logarithm)
- 📐 **Trigonometric Functions**:  
  `sin`, `cos`, `tan`, `arcsin` (Inverse sine), `arccos` (Inverse sine), `arctan` (Inverse tangent) 
- 🔒 **Logical Operations**:  
  `and`, `&` (AND), `or`, `|` (OR), `xor` (XOR), `^` (Bitwise XOR), `not` (Bitwise NOT)
- 🎩 **Factorial**:  
  `!` (Factorial of a number), `fact`
- 🎨 **Styled Output**:  
  Beautiful and clear output using the `lipgloss` library. 🌈

---

### Installation
```bash
   git clone https://github.com/your-username/rpn-calculator.git
   cd rpn-calculator
   go mod tidy
   go run . <expression>
```

---
## ✨ Examples
Basic Arithmetic
```bash
  go run main.go 3 5 +
# Output: 8
```
Trigonometry
```bash
  go run main.go 3.14159 cos
  # Output: -1
  
  go run main.go 0.5 sin
  # Output: 0.479425538604203

```
Logical Operations
```bash
  go run main.go 5 3 and
  # Output: 1
  
  go run main.go 5 3 or
  # Output: 7
  
  go run main.go 5 3 xor
  # Output: 6
  8
```
Advanced Operations
```bash
  go run main.go 9 sqrt
  # Output: 3
  
  go run main.go 2 3 pow
  # Output: 8
  
  go run main.go 5 fact
  # Output: 120
```
More Complex Examples
```bash
  # Complex expression with multiple operations
  go run main.go 5 3 8 x 7 + - 
  # Output: -26
  
  # Using multiple functions
  go run main.go 3.14159 2 pow 2 sqrt + sin
  # Output: -0.9587446221046463 
```














