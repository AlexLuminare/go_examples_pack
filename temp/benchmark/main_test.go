package main

import (
	"testing"
)

func BenchmarkSwapValuesP(b *testing.B) {
	for range b.N {
		swapValuesP(&a1, &a2)
	}
}

func BenchmarkSwapValuesV(b *testing.B) {
	for range b.N {
		swapValuesV(a1, a2)
	}
}
