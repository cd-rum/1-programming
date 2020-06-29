package main

import "fmt"

func main() {
  fmt.Println(factorial(6))
}

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return n * factorial(n - 1)
}

