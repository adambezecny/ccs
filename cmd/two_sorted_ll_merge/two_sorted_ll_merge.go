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

func MergeLists(list1 []int, list2 []int) []int {
	list1 = append(list1, list2...)
	sort.Ints(list1)
	return list1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int(tTemp)

	for ii := 0; ii < t; ii++ {
		var list1 []int
		var list2 []int

		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int(nTemp)

		for i := 0; i < n; i++ {
			item, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
			checkError(err)
			list1 = append(list1, int(item))
		}

		mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		m := int(mTemp)

		for i := 0; i < m; i++ {
			item, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
			checkError(err)
			list2 = append(list2, int(item))
		}

		result := MergeLists(list1, list2)
		for _, item := range result {
			fmt.Printf("%d ", item)
		}
		fmt.Println("")
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
