package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPairs(t *testing.T) {
	assert.Equal(t, int32(3), pairsSlow(1, []int32{1, 2, 3, 4}))
	assert.Equal(t, int32(3), pairs(1, []int32{1, 2, 3, 4}))
}
