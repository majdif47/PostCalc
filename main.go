package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: pcalc <expression>")
		return
	}

	args := os.Args[1:]
  resu, exp, err := evaluate(args)
  if err != nil {
    fmt.Println("Error")
    return
  }
	// Final result
  expression := styledBox("Expression:\t"+fmt.Sprint(exp), "#AD88C6")
  fmt.Println(expression)

  result := styledBox("Result:\t"+fmt.Sprint(resu), "#77B0AA")
  fmt.Println(result)  
}

