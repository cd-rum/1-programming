import time

fn main() {
  sw := time.new_stopwatch({})
  println(fact_one(6))
  println(sw.elapsed().nanoseconds())
  println(fact_two(6))
  println(sw.elapsed().nanoseconds())
  println(fib_one(7))
  println(sw.elapsed().nanoseconds())
  println(fib_two(7))
  println(sw.elapsed().nanoseconds())
  println(count_change(100))
  println(sw.elapsed().nanoseconds())
  println(tree_recursive(100))
  println(sw.elapsed().nanoseconds())
}

fn fact_one(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact_one(n - 1)
}

fn fact_two(n int) int {
  return fact_iter(1, 1, n)
}

fn fact_iter(product int, counter int, max_count int) int {
  if counter > max_count {
    return product
  }
  return fact_iter(product * counter, counter + 1, max_count)
}

fn fib_one(n int) int {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  }
  return fib_one(n - 1) + fib_one(n - 2)
}

fn fib_two(n int) int {
  return fib_iter(1, 0, n)
}

fn fib_iter(a int, b int, count int) int {
  if count == 0 {
    return b
  }
  return fib_iter(a + b, a, count - 1)
}

fn count_change(amount int) int {
  return cc(amount, 5)
}

fn cc(amount int, kinds_of_coins int) int {
  if amount == 0 {
    return 1
  }  else if amount < 0 || kinds_of_coins == 0 {
    return 0
  }
  return cc(amount, kinds_of_coins - 1) + cc(amount - first_denomination(kinds_of_coins), kinds_of_coins)
}

// needs to be exhaustive
fn first_denomination(kinds_of_coins int) int {
  return match kinds_of_coins {
    1    { 1 }
    2    { 5 }
    3    { 10 }
    4    { 25 }
    5    { 50 }
    else { 50 }
  }
}

fn tree_recursive(n int) int {
  if n < 3 {
    return n
  }
  return tree_iterative(2, 1, 0, n)
}

fn tree_iterative(a int, b int, c int, count int) int {
  if count < 3 {
    return a
  }
  return tree_iterative(a + (2 * b) + (3 * c), a, b, count - 1)
}

fn triangle(depth u16) [][]u64 {
  mut tri := []u64{ cap: int(depth), init: 1 }

  for d, _ in tri {
    mut row := []u64{ cap: int(d + 1) }
    for col, _ in d / 2 {
      val := u64(1)
      if col > 0 {
        prev := tri[d-1]
        val = prev[col-1] + prev[col]
      }
      row[col], row[d-col] = val, val
    }
    tri[d] = row
  }
  return tri
}
