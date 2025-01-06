package main

import (
  "fmt"
  "strconv"
  stack "github.com/majdif47/Go-dsa/ds/stacks"
)

func evaluate(args []string) (float64,string,error) {
	s := stack.NewStack[float64]()
  exp := "("
	for _, arg := range args {
    exp += " "
    exp += arg
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
      }else if constant, in := constants[arg]; in{
        s.Push(constant)
      }else {
				return 0, "", fmt.Errorf("Error: Unknown operator", arg)
				
			}
		}
	}
  if s.Size() != 1 {
		return 0, "", fmt.Errorf("Invalid expression!")
	}
  exp += " )"
  result, _ := s.Pop()
  return result, exp,nil
}
