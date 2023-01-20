package main

import (
  "fmt"
  "math/rand"
)

func main() {
  for i := 0; i < 10; i++ {
    fmt.Printf("%d) %d\n", i, rand.Intn(25))
  }
}