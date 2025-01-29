package main

import (
	"math/rand/v2"
	"slices"
)

func getRandom(n int) []int {
	result := make([]int, 0, n)

	for {
		num := rand.IntN(45) + 1
		if !slices.Contains(result, num) {
			result = append(result, num)
		}
		if len(result) == n {
			break
		}
	}
	slices.Sort(result)
	return result
}
