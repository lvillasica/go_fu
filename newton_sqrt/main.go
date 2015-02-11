package main

import(
  "fmt"
  "github.com/pkgs/newtonSqrt"
)

const val = 4

func main() {
  sqrt := newtonSqrt.Sqrt(val)
  fmt.Println("Final square root of ", val, ": ", sqrt)
}