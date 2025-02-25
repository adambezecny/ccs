package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tree := map[int32]Node{
		1: {Value: 1, Neighbors: []int32{2, 3}},
		2: {Value: 2, Neighbors: []int32{}},
		3: {Value: 3, Neighbors: []int32{4}},
		4: {Value: 4, Neighbors: []int32{}},
		5: {Value: 5, Neighbors: []int32{}},
	}
	var result string

	result = breadthFirstSearch(tree, 1, 2)
	assert.Equal(t, "1->2", result)
	assert.Equal(t, int32(6), pathNumber(result))

	result = breadthFirstSearch(tree, 1, 3)
	assert.Equal(t, "1->3", result)
	assert.Equal(t, int32(6), pathNumber(result))

	result = breadthFirstSearch(tree, 1, 4)
	assert.Equal(t, "1->3->4", result)
	assert.Equal(t, int32(12), pathNumber(result))

	result = breadthFirstSearch(tree, 1, 5)
	assert.Equal(t, "not_found", result)
	assert.Equal(t, int32(-1), pathNumber(result))
}

func TestBfs1(t *testing.T) {
	edges := [][]int32{
		{1, 2},
		{1, 3},
		{3, 4},
	}
	result := bfs(int32(5), int32(3), edges, int32(1))
	assert.Equal(t, []int32{6, 6, 12, -1}, result)
}

func TestBfs2(t *testing.T) {
	edges := [][]int32{
		{1, 2},
		{1, 3},
	}
	result := bfs(int32(4), int32(2), edges, int32(1))
	assert.Equal(t, []int32{6, 6, -1}, result)
}

func TestBfs3(t *testing.T) {
	edges := [][]int32{
		{2, 3},
	}
	result := bfs(int32(3), int32(1), edges, int32(2))
	assert.Equal(t, []int32{-1, 6}, result)
}
