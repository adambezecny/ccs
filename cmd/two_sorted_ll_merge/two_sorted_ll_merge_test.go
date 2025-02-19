package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeList(t *testing.T) {
	assert.Equal(t,
		[]int{1, 2, 3, 3, 4},
		MergeLists([]int{1, 2, 3}, []int{3, 4}),
	)
}
