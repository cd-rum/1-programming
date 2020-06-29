package main

import (
    "testing"
)

func BenchmarkFactOne(b *testing.B) {
    for i := 0; i < b.N; i++ {
        factOne(6)
    }
}

func BenchmarkFactTwo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        factTwo(6)
    }
}
