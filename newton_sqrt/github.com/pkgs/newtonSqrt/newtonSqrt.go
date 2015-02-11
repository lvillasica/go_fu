package newtonSqrt

import "fmt"

func Sqrt(x float64) float64 {
  z := 1.0
  for steps := 1; steps <= 10; steps++ {
    z = newtonSqrt(x, z)
    fmt.Println(z)
  }
  return z
}

func newtonSqrt(x, z float64) float64 {
  z -= (((z*z) - x) / (2 * z))
  return z
}