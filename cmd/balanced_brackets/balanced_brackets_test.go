package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	assert.Equal(t, "YES", isBalanced("{[()]}"))
	assert.Equal(t, "NO", isBalanced("{[(])}"))
	assert.Equal(t, "YES", isBalanced("{{[[(())]]}}"))
}
