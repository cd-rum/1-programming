package main

import "fmt"

func main() {
  fmt.Println(factOne(6))
  fmt.Println(factTwo(6))
}

func factOne(n int) int {
  if n == 0 {
    return 1
  }
  return n * factOne(n - 1)
}

func factTwo(n int) int {
  return factIter(1, 1, n)
}

func factIter(product int, counter int, maxCount int) int {
  if counter > maxCount {
    return product
  }
  return factIter(product * counter, counter + 1, maxCount)
}
