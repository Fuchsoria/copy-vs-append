package main

import (
	"math/rand"
)

var sourceArr100 []int = createSourceArray(100)
var sourceArr100k []int = createSourceArray(100_000)
var sourceArr100kk []int = createSourceArray(100_000_000)

func createSourceArray(n int) (arr []int) {
	arr = make([]int, n)
	min := 0
	max := n

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}

	return arr
}
