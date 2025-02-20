package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLegoBlocks(t *testing.T) {
	assert.Equal(t, int32(0), legoBlocks(2, 3))
}
