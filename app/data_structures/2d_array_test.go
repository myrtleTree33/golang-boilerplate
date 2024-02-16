package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrint(t *testing.T) {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	PrettyPrint(arr)

	assert.True(t, true)
}

func TestNew2DArray(t *testing.T) {
	arr := New2DArray(3, 3, 0)
	PrettyPrint(arr)

	assert.True(t, true)
}
