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
	elements []string
}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds an element to the stack
func (s *Stack) Push(value string) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() string {
	if len(s.elements) == 0 {
		return "-1"
	}
	// Get the last element
	topElement := s.elements[len(s.elements)-1]
	// Remove the last element
	s.elements = s.elements[:len(s.elements)-1]
	return topElement
}

// Peek returns the top element without removing it
func (s *Stack) Peek() string {
	if len(s.elements) == 0 {
		return "-1"
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func isBalanced(s string) string {
	brackets := strings.Split(s, "")

	stack := NewStack()

	for _, bracket := range brackets {
		if bracket == "{" || bracket == "[" || bracket == "(" {
			stack.Push(bracket)
		} else /* is ')' or ']' or '}' */ {
			popped := stack.Pop()
			if popped == "-1" {
				return "NO"
			}

			if (popped == "{" && bracket == "}") ||
				(popped == "(" && bracket == ")") ||
				(popped == "[" && bracket == "]") {
				continue
			}
			return "NO"
		}
	}

	if stack.IsEmpty() {
		return "YES"
	}
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		s := readLine(reader)

		result := isBalanced(s)

		fmt.Printf("%s\n", result)
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
