package main

import "fmt"

func main() {
  fmt.Println(factOne(6))
  fmt.Println(factTwo(6))

  fmt.Println(fibOne(7))
  fmt.Println(fibTwo(7))

  fmt.Println(countChange(100))
  fmt.Println(treeRecursive(100))
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

func countChange(amount int) int {
  return cc(amount, 5)
}

func cc(amount int, kindsOfCoins int) int {
  if amount == 0 {
    return 1
  } else if amount < 0 || kindsOfCoins == 0 {
    return 0
  }
  return cc(amount, kindsOfCoins - 1) + cc(amount - firstDenomination(kindsOfCoins), kindsOfCoins)
}

func firstDenomination(kindsOfCoins int) int {
  switch kindsOfCoins {
    case 1: return 1
    case 2: return 5
    case 3: return 10
    case 4: return 25
    case 5: return 50
  }
  return 1
}

func treeRecursive(n int) int {
  if n < 3 {
    return n
  }
  return treeIterative(2, 1, 0, n)
}

func treeIterative(a int, b int, c int, count int) int {
  if count < 3 {
    return a
  }
  return treeIterative(a + (2 * b) + (3 * c), a, b, count - 1)
}
