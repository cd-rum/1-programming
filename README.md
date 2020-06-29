# Structure and interpretation of computer programs

Completing [this book](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html) in Go.

```zsh
go test -bench=.
```

Lower is [better](https://golang.org/pkg/testing/).

> The acts of the mind, wherein it exerts its power over simple ideas, are chiefly these three: 1. Combining several simple ideas into one compound one, and thus all complex ideas are made. 2. The second is bringing two ideas, whether simple or complex, together, and setting them by one another so as to take a view of them at once, without uniting them into one, by which it gets all its ideas of relations. 3. The third is separating them from all other ideas that accompany them in their real existence: this is called abstraction, and thus all its general ideas are made.
> - _John Locke, An Essay Concerning Human Understanding (1690)_

The ability to visualize the consequences of the actions under consideration is crucial to becoming an expert programmer, just as it is in any synthetic, creative activity.

### Linear recursion and iteration

We can compute `n!` by computing `(n - 1)!` and multiplying the result by `n`. If we add the stipulation that `1!` is equal to `1`, this translates into a procedure:

```lisp
(define (factorial n)
  (if (= n 1)
      1
      (* n (factorial (- n 1)))))
```

Now, other approach. We could describe a rule for computing `n!` by specifying that we first multiply `1` by `2`, then multiply the result by `3`, then by `4`, and so on until we reach `n`. More formally, we maintain a *running product*, together with a counter that counts from `1` up to `n`. Or:

```
product = counter * product
counter = counter + 1
```

Or in Lisp:

```lisp
(define (factorial n)
  (fact-iter 1 1 n))

(define (fact-iter product counter max-count)
  (if (> counter max-count)
      product
      (fact-iter (* counter product)
                 (+ counter 1)
                 max-count)))
```

When we consider the _shapes_ of the two processes, we find that they evolve quite differently. Consider the first process. The substitution model reveals a shape of expansion followed by contraction. The expansion occurs as the process builds up a chain of deferred operations (in this case, a chain of multiplications). ... **Carrying out this process requires that the interpreter keep track of the operations to be performed later on.**

By contrast, the second process does not grow and shrink. At each step, all we need to keep track of, for any `n`, are the current values of the variables `product`, `counter`, and `max-count`. We call this an iterative process. In general, an iterative process is one whose state can be summarized by a fixed number of state variables, together with a fixed rule that describes how the state variables should be updated as the process moves from state to state and an (optional) end test that specifies conditions under which the process should terminate.

**In the iterative case, the program variables provide a complete description of the state of the process at any point.**

### Tree recursion

```lisp
(define (fib n)
  (cond ((= n 0) 0)
        ((= n 1) 1)
        (else (+ (fib (- n 1))
                 (fib (- n 2))))))
```

**This procedure is instructive as a prototypical tree recursion, but it is a terrible way to compute Fibonacci numbers because it does so much redundant computation.**

Given:

```
a = a + b
b = a
```

It is not hard to show that, after applying this transformation `n` times, `a` and `b` will be equal.

```lisp
(define (fib n)
  (fib-iter 1 0 n))

(define (fib-iter a b count)
  (if (= count 0)
      b
      (fib-iter (+ a b) a (- count 1))))
```

The Go equivalents perform now thus:

```golang
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cd-rum/programming
BenchmarkFactOne-16     88593147                13.4 ns/op
BenchmarkFactTwo-16     76278255                16.2 ns/op
BenchmarkFibOne-16      16317018                73.5 ns/op
BenchmarkFibTwo-16      68011566                18.0 ns/op
PASS
ok      github.com/cd-rum/programming   6.129s
```

It takes only a bit of cleverness to come up with the iterative Fibonacci algorithm. In contrast, consider the following problem: How many different ways can we make change of $1.00, given half-dollars, quarters, dimes, nickels, and pennies? More generally, can we write a procedure to compute the number of ways to change any given amount of money?

 This problem has a simple solution as a recursive procedure. Suppose we think of the types of coins available as arranged in some order. Then the following relation holds:

The number of ways to change amount `a` using `n` kinds of coins equals:
* the number of ways to change amount `a` using all but the first kind of coin, plus
* the number of ways to change amount `a - d` using all `n` kinds of coins, where `d` is the denomination of the first kind of coin.

**To see why this is true, observe that the ways to make change can be divided into two groups: those that do not use any of the first kind of coin, and those that do. Therefore, the total number of ways to make change for some amount is equal to the number of ways to make change for the amount without using any of the first kind of coin, plus the number of ways to make change assuming that we do use the first kind of coin. But the latter number is equal to the number of ways to make change for the amount that remains after using a coin of the first kind.**

...

```lisp
(define (count-change amount)
  (cc amount 5))
(define (cc amount kinds-of-coins)
  (cond ((= amount 0) 1)
        ((or (< amount 0) (= kinds-of-coins 0)) 0)
        (else (+ (cc amount
                     (- kinds-of-coins 1))
                 (cc (- amount
                        (first-denomination kinds-of-coins))
                     kinds-of-coins)))))
(define (first-denomination kinds-of-coins)
  (cond ((= kinds-of-coins 1) 1)
        ((= kinds-of-coins 2) 5)
        ((= kinds-of-coins 3) 10)
        ((= kinds-of-coins 4) 25)
        ((= kinds-of-coins 5) 50)))
```

Intuitively, this felt off to me but it's about first denominations and not amounts.

```zsh
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/cd-rum/programming
BenchmarkFactOne-16             84137239                14.0 ns/op
BenchmarkFactTwo-16             75466722                16.3 ns/op
BenchmarkFibOne-16              15745932                76.9 ns/op
BenchmarkFibTwo-16              64720564                18.2 ns/op
BenchmarkCountChange-16            26160             44524 ns/op
PASS
ok      github.com/cd-rum/programming   9.319s
```

It ain't fast. The authors (Abelson, Sussman, and Sussman) write that `countChange` generates a tree-recursive process with redundancies similar to those in our first implementation of `fib`. The observation that a tree-recursive process may be highly inefficient but often easy to specify and understand has led people to propose that one could get the best of both worlds by designing a _smart compiler_ that could transform tree-recursive procedures into more efficient procedures that compute the same result. Easy to understand but slow.
