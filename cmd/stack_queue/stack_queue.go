package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Stack struct using slice
type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds an element to the stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	// Get the last element
	topElement := s.elements[len(s.elements)-1]
	// Remove the last element
	s.elements = s.elements[:len(s.elements)-1]
	return topElement
}

// Peek returns the top element without removing it
func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		return -1
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Queue struct using slice
type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue adds an element to the queue
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element from the queue
func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		return -1
	}
	// Get the first element
	frontElement := q.elements[0]
	// Remove the first element
	q.elements = q.elements[1:]
	return frontElement
}

// Peek returns the front element without removing it
func (q *Queue) Peek() int {
	if len(q.elements) == 0 {
		return -1
	}
	return q.elements[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

/*
queue using two stacks
https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0232.implement-queue-using-stacks/src/queue-by-stacks.go
https://stackoverflow.com/questions/77814807/creating-a-queue-using-two-stack-but-make-enqueuing-o1
*/

// StackQueue defines Stack1 queue by two stacks
type StackQueue struct {
	Stack1, Stack2 *stack
}

func NewStackQueue() StackQueue {
	return StackQueue{
		Stack1: newStack(),
		Stack2: newStack(),
	}
}

// Push element x to the back of queue.
func (queue *StackQueue) Push(x int) {
	queue.Stack1.push(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (queue *StackQueue) Pop() int {
	if queue.Stack2.isEmpty() {
		for queue.Stack1.len() > 1 {
			queue.Stack2.push(queue.Stack1.pop())
		}
		return queue.Stack1.pop()
	}
	return queue.Stack2.pop()
}

// Peek Get the front element.
func (queue *StackQueue) Peek() int {
	res := queue.Pop()
	if res != -1 {
		queue.Stack2.push(res)
	}
	return res
}

// Empty Returns whether the queue is empty.
func (queue *StackQueue) Empty() bool {
	return queue.Stack1.isEmpty() && queue.Stack2.isEmpty()
}

// stack defines Stack1 stack
type stack struct {
	nums []int
}

// newStack creates a empty stack
func newStack() *stack {
	return &stack{
		nums: []int{},
	}
}

func (s *stack) push(n int) {
	s.nums = append(s.nums, n)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *stack) len() int {
	return len(s.nums)
}
func (s *stack) isEmpty() bool {
	return s.len() == 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int(qTemp)

	queue := NewQueue()

	var operations [][]string

	for i := 0; i < q; i++ {
		operation := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		operations = append(operations, operation)

	}

	for _, operation := range operations {
		operationType, err := strconv.ParseInt(operation[0], 10, 64)
		checkError(err)

		if operationType == 1 /* enqueue */ {
			valStr := operation[1]
			val, err := strconv.ParseInt(valStr, 10, 64)
			checkError(err)
			queue.Enqueue(int(val))
		}

		if operationType == 2 /* dequeue */ {
			queue.Dequeue()
		}

		if operationType == 3 /* print */ {
			fmt.Printf("%d\n", queue.Peek())
		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
