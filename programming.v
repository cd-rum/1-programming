import time

fn main() {
  sw := time.new_stopwatch({})
  println(fact_one(6))
  println(sw.elapsed().nanoseconds())
  println(fact_two(6))
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
