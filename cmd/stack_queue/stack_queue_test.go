package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	assert.Equal(t, -1, stack.Peek())
	assert.Equal(t, -1, stack.Pop())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(t, 3, stack.Peek())
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, -1, stack.Pop())

}

func TestQueue(t *testing.T) {
	queue := NewQueue()
	assert.Equal(t, -1, queue.Peek())
	assert.Equal(t, -1, queue.Dequeue())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Dequeue())
	assert.Equal(t, 2, queue.Dequeue())
	assert.Equal(t, 3, queue.Dequeue())
	assert.Equal(t, -1, queue.Dequeue())
}

func TestStackQueue(t *testing.T) {
	queue := NewStackQueue()
	assert.Equal(t, -1, queue.Peek())
	assert.Equal(t, -1, queue.Pop())

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Pop())
	assert.Equal(t, 2, queue.Pop())
	assert.Equal(t, 3, queue.Pop())
	assert.Equal(t, -1, queue.Pop())
}
