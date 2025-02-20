package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCookies(t *testing.T) {
	assert.Equal(t, int32(4), cookies(9, []int32{2, 7, 3, 6, 4, 6}))
	assert.Equal(t, int32(2), cookies(7, []int32{1, 2, 3, 9, 10, 12}))
	assert.Equal(t, int32(-1), cookies(1000, []int32{1, 2, 3}))
	assert.Equal(t, int32(0), cookies(1, []int32{2, 3, 4}))
}
