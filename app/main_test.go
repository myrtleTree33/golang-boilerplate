package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse[T comparable](arr []T) []T {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return arr
}

func TestMain(t *testing.T) {

	a := []int{1, 2, 3, 4, 5}
	t.Log(reverse(a))

	assert.Equal(t, 2, 2)
}
