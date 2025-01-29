package main

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandom(t *testing.T) {
	assert := assert.New(t)
	for range 5 {
		ran := rand.IntN(10) + 1
		result := getRandom(ran)
		assert.Equal(ran, len(result), "expected %d elements, got %d", ran, len(result))
		assert.True(slices.IsSorted(result), "result is not sorted")
		assert.True(ran == len(slices.Compact(result)), "result has duplicates")
	}
}
