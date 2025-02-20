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
	elements [][]string
}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds an element to the stack
func (s *Stack) Push(value []string) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() []string {
	if len(s.elements) == 0 {
		return nil
	}
	// Get the last element
	topElement := s.elements[len(s.elements)-1]
	// Remove the last element
	s.elements = s.elements[:len(s.elements)-1]
	return topElement
}

// Peek returns the top element without removing it
func (s *Stack) Peek() []string {
	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func executeOperations(content []string, operations [][]string) []string {
	var undoableOperations = NewStack()
	output := make([]string, 0)

	// operation[0] - type, operation[1] - argument
	for _, operation := range operations {
		operationType, err := strconv.ParseInt(operation[0], 10, 64)
		checkError(err)

		if operationType == 1 /* append */ {
			undoableOperations.Push(operation)
			content = append(content, strings.Split(operation[1], "")...)
			continue
		}

		if operationType == 2 /* delete */ {
			k, err := strconv.ParseInt(operation[1], 10, 64)
			checkError(err)

			if len(content) >= int(k) {
				removed := content[len(content)-int(k):]
				undoableOperations.Push(append(operation, strings.Join(removed, "")))
				content = content[:len(content)-int(k)]
			}
			continue
		}

		if operationType == 3 /* print */ {
			k, err := strconv.ParseInt(operation[1], 10, 64)
			checkError(err)
			if k > 0 && int(k) <= len(content) {
				//fmt.Println(content[k-1 : k])
				output = append(output, strings.Join(content[k-1:k], ""))
			}
			continue
		}

		if operationType == 4 /* undo */ {
			undoOper := undoableOperations.Pop()
			if undoOper != nil {
				if undoOper[0] == "1" /* append -> reverse operation is delete */ {
					// remove last k chars where k is length of appended data
					k := len(undoOper[1])
					content = content[:len(content)-k]
				}

				if undoOper[0] == "2" /* delete -> reverse operation is add*/ {
					content = append(content, strings.Split(undoOper[2], "")...) // add what was deleted
				}
			}
			continue
		}
	}

	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int(qTemp)

	var operations [][]string

	for i := 0; i < q; i++ {
		operation := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		operations = append(operations, operation)

	}

	result := executeOperations(make([]string, 0), operations)
	for _, item := range result {
		fmt.Println(item)
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
