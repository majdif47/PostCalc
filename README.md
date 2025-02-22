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
   git clone https://github.com/majdif47/PostCalc.git && cd PostCalc && go mod tidy

```
```bash
  go build -o pcal
```
---
## ✨ Examples
Basic Arithmetic
```bash
  ./pcal 3 5 +
# Output: 8
```
Trigonometry
```bash
  ./pcal 3.14159 cos
  # Output: -1
  
  ./pcal 0.5 sin
  # Output: 0.479425538604203

```
Logical Operations
```bash
  ./pcal 5 3 and
  # Output: 1
  
  ./pcal 5 3 or
  # Output: 7
  
  ./pcal 5 3 xor
  # Output: 6
  8
```
Advanced Operations
```bash
  ./pcal 9 sqrt
  # Output: 3
  
  ./pcal 2 3 pow
  # Output: 8
  
  ./pcal 5 fact
  # Output: 120
```
More Complex Examples
```bash
  # Complex expression with multiple operations
  ./pcal 5 3 8 x 7 + - 
  # Output: -26
  
  # Using multiple functions
  ./pcal 3.14159 2 pow 2 sqrt + sin
  # Output: -0.9587446221046463 
```














