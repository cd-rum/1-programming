package main

import "fmt"

func main() {
  fmt.Println(factOne(6))
  fmt.Println(factTwo(6))

  fmt.Println(fibOne(7))
  fmt.Println(fibTwo(7))
}

/*
  Examples are written to maintain logical similarity to the lisp examples,
  not to be idiomatic Go. That would likely ruin the point of the exercises,
  too,
*/

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

func fibOne(n int) int {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  }
  return fibOne(n - 1) + fibOne(n - 2)
}

func fibTwo(n int) int {
  return fibIter(1, 0, n)
}

func fibIter(a int, b int, count int) int {
  if count == 0 {
    return b
  }
  return fibIter(a + b, a, count - 1)
}
