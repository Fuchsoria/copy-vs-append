package main

import (
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	b.Run("append 100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var appendArr []int

			_ = append(appendArr, sourceArr100...)
		}
	})

	b.Run("append 100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var appendArr []int

			_ = append(appendArr, sourceArr100k...)
		}
	})

	b.Run("append 100kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var appendArr []int

			_ = append(appendArr, sourceArr100kk...)
		}
	})
}

func BenchmarkCopy(b *testing.B) {
	b.Run("copy 100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copiedArr := make([]int, len(sourceArr100))

			_ = copy(copiedArr, sourceArr100)
		}
	})

	b.Run("copy 100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copiedArr := make([]int, len(sourceArr100k))

			_ = copy(copiedArr, sourceArr100k)
		}
	})

	b.Run("copy 100kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copiedArr := make([]int, len(sourceArr100kk))

			_ = copy(copiedArr, sourceArr100kk)
		}
	})
}
