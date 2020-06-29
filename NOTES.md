# Structure and interpretation of computer programs

[Required text](https://mitpress.mit.edu/sites/default/files/sicp/full-text/book/book.html)

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

When we consider the _shapes_ of the two processes, we find that they evolve quite differently. Consider the first process. The substitution model reveals a shape of expansion followed by contraction. The expansion occurs as the process builds up a chain of deferred operations (in this case, a chain of multiplications).
