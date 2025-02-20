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

/*
 * Complete the 'cookies' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 */

func cookies(k int32, arr []int32) int32 {
	var iterations int32 = 0

	sweetEnough := func(k int32, cookies []int) bool {
		for _, cookie := range cookies {
			if cookie < int(k) {
				return false
			}
		}

		return true
	}

	sortedCookies := make([]int, len(arr))
	for idx, item := range arr {
		sortedCookies[idx] = int(item)
	}
	sort.Ints(sortedCookies)

	if sweetEnough(k, sortedCookies) {
		return 0
	}

	found := false
	for {
		iterations++
		twoLeftMostCookies := sortedCookies[0:2]
		sortedCookies = sortedCookies[2:]

		mixedCookie := twoLeftMostCookies[0] + 2*twoLeftMostCookies[1]
		sortedCookies = append(sortedCookies, mixedCookie)
		sort.Ints(sortedCookies)

		if sweetEnough(k, sortedCookies) {
			found = true
			break
		}

		if len(sortedCookies) == 1 {
			break
		}
	}

	if found {
		return iterations
	} else {
		return -1
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var A []int32

	for i := 0; i < int(n); i++ {
		AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
		checkError(err)
		AItem := int32(AItemTemp)
		A = append(A, AItem)
	}

	result := cookies(k, A)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
