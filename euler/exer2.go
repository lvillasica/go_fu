package main

import (
  "fmt"
  "math"
)

type FibonacciSeq struct {
  prev1, prev2, curr int
}

func (seq FibonacciSeq) evenCurr() (val int) {
  if math.Mod(float64(seq.curr), 2.0) == 0 {
    val = seq.curr
  }
  return
}

// https://projecteuler.net/problem=2
func main() {
  fmt.Println("Total of even fibonacci sets below 4M: ", totalEvenFibonacci(4000000))
}

func totalEvenFibonacci(limit int) (total int) {
  seq := FibonacciSeq{0, 0, 1}
  for seq.curr < limit {
    total += seq.evenCurr()
    seq.prev1 = seq.prev2
    seq.prev2 = seq.curr
    if seq.curr > 2 {
      seq.curr = seq.prev1 + seq.prev2
    } else {
      seq.curr += 1
    }
  }
  return
}