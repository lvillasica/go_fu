package main

import (
  "fmt"
  "math"
)

type Number struct {
  num int
}

func (n Number) isMultipleOf(m int) bool {
  return math.Mod(float64(n.num), float64(m)) == 0
}

// https://projecteuler.net/problem=1
func main() {
  fmt.Println("Total of all the multiples of 3 or 5 below 1000: ", totalByMultOf3Or5(1000))
}

func totalByMultOf3Or5(limit int) (total int) {
  for i := 1; i < limit; i++ {
    n := Number{i}
    if n.isMultipleOf(3) || n.isMultipleOf(5) {
      total += n.num
    }
  }
  return
}