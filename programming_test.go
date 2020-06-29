package main

import (
    "testing"
)

func BenchmarkFactorialOne(b *testing.B) {
    for i := 0; i < b.N; i++ {
        factorialOne(6)
    }
}

func BenchmarkFactorialTwo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        factorialTwo(6)
    }
}
