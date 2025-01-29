package main

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetRandom(t *testing.T) {
	for range 5 {
		ran := rand.IntN(10) + 1
		result := getRandom(ran)
		assert.Equal(t, ran, len(result))
		assert.Equal(t, ran, len(slices.Compact(result)))
	}
}
