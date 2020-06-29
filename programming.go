package main

import "fmt"

func main() {
  fmt.Println(factorialOne(6))
  fmt.Println(factorialTwo(6))
}

func factorialOne(n int) int {
  if n == 0 {
    return 1
  }
  return n * factorialOne(n - 1)
}

func factorialTwo(n int) int {
  return factIter(1, 1, n)
}

func factIter(product int, counter int, n int) int {
  if counter > n {
    return product
  }
  return factIter(product * counter, counter + 1, n)
}
