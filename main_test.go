package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("copy", func(t *testing.T) {
		copiedArr := make([]int, len(sourceArr100))
		copy(copiedArr, sourceArr100)

		require.Equal(t, sourceArr100, copiedArr)
	})

	t.Run("copy immutable", func(t *testing.T) {
		copiedArr := make([]int, len(sourceArr100))
		copy(copiedArr, sourceArr100)

		require.Len(t, copiedArr, len(sourceArr100))

		copiedArr[42] = 4242 // generated from 1 to 100, 4242 is new number

		require.NotEqual(t, sourceArr100[42], copiedArr[42])
	})
}

func TestAppend(t *testing.T) {
	t.Run("append", func(t *testing.T) {
		var appenedArr []int
		appenedArr = append(appenedArr, sourceArr100...)

		require.Equal(t, sourceArr100, appenedArr)
	})

	t.Run("append immutable", func(t *testing.T) {
		var appenedArr []int
		appenedArr = append(appenedArr, sourceArr100...)

		require.Len(t, appenedArr, len(sourceArr100))

		appenedArr[42] = 4242 // generated from 1 to 100, 4242 is new number

		require.NotEqual(t, sourceArr100[42], appenedArr[42])
	})
}
