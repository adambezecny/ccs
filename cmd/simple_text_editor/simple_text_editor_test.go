package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	assert.Nil(t, stack.Peek())
	assert.Nil(t, stack.Pop())

	stack.Push([]string{"1"})
	stack.Push([]string{"2"})
	stack.Push([]string{"3"})
	assert.Equal(t, []string{"3"}, stack.Peek())
	assert.Equal(t, []string{"3"}, stack.Pop())
	assert.Equal(t, []string{"2"}, stack.Pop())
	assert.Equal(t, []string{"1"}, stack.Pop())
	assert.Nil(t, stack.Pop())
}

func TestRemovelastK(t *testing.T) {
	str := "abcde"

	//remove last two
	assert.Equal(t, "abc", str[:len(str)-2])
	// removed
	assert.Equal(t, "de", str[len(str)-2:])
}

func TestExecuteOperations(t *testing.T) {
	initialString := []string{"a", "b", "c", "d", "e"}

	operations := [][]string{
		{"1", "fg"},
		{"3", "6"},
		{"2", "5"},
		{"4"},
		{"3", "7"},
		{"4"},
		{"3", "4"},
	}

	result := executeOperations(initialString, operations)
	assert.Equal(t, []string{"f", "g", "d"}, result)
}
