package main

import "math"



const epsilon = 1e-15

func fact(num int) int {
	if num <= 1 {
		return 1
	}
	return num * fact(num-1)
}

func roundToZero(num float64) float64 {
  if math.Abs(num) < epsilon {
    return 0
  }
  return num
}
