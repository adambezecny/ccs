package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func findZigZagSequence(arr []int64) []int64 {
	// sort in ascending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	middleIdx := (len(arr)+1)/2 - 1

	// swap the biggest and middle elements
	swap := arr[middleIdx]
	arr[middleIdx] = arr[len(arr)-1]
	arr[len(arr)-1] = swap

	descendingElements := arr[middleIdx+1:]

	// sort elements in descending order
	sort.Slice(descendingElements, func(i, j int) bool {
		return descendingElements[i] > descendingElements[j]
	})

	// replace them in original array
	copy(arr[middleIdx+1:], descendingElements)

	return arr
}

/*
	func main() {
		reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

		stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
		checkError(err)

		defer stdout.Close()
		writer := bufio.NewWriterSize(stdout, 16*1024*1024)

		t, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var a []int64

		for i := 0; i < int(t); i++ {
			for i := 0; i < int(n); i++ {
				aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
				checkError(err)
				aItem := aItemTemp
				a = append(a, aItem)
			}

			result := findZigZagSequence(a)

			for idx, resultItem := range result {
				fmt.Fprintf(writer, "%d", resultItem)

				if idx != len(result)-1 {
					fmt.Fprintf(writer, " ")
				}
			}

			fmt.Fprintf(writer, "\n")
			writer.Flush()
		}
	}
*/
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	t, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int64

	for i := 0; i < int(t); i++ {
		for i := 0; i < int(n); i++ {
			aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			aItem := aItemTemp
			a = append(a, aItem)
		}

		result := findZigZagSequence(a)

		for idx, resultItem := range result {
			fmt.Printf("%d", resultItem)

			if idx != len(result)-1 {
				fmt.Printf(" ")
			}
		}

		fmt.Printf("\n")
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
