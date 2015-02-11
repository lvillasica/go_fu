package main

import (
  "fmt"
  "math"
  "math/big"
)

type Number struct {
  num int64
}

func (factor Number) isPrimeFactorFor(n Number) bool {
  bigFactor := big.NewInt(factor.num)
  return bigFactor.ProbablyPrime(1) == true && math.Mod(float64(n.num), float64(factor.num)) == 0
}

func (n Number) largestPrimeFactor() (largestPrime int64) {
  factor := Number{2}
  for n.num > 1 {
    if factor.isPrimeFactorFor(n) {
      n.num = n.num / factor.num
      if factor.num > largestPrime {
        largestPrime = factor.num
      }
    } else {
      factor.num += 1
    }
  }
  return largestPrime
}

// https://projecteuler.net/problem=3
func main() {
  num := Number{600851475143}
  fmt.Println("largest prime factor of the number 600851475143: ", num.largestPrimeFactor())
}