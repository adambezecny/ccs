package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

func minimumBribes(q []int32) {
	counter := 0
	for idx, val := range q {
		normalPos := int(val - 1)
		realPos := idx
		diff := normalPos - realPos

		if diff > 0 && diff <= 2 {
			counter += diff
		}

		if diff > 0 && diff > 2 {
			fmt.Println("Too chaotic")
			return
		}
	}

	fmt.Println(counter)
}

/*
2 1 5 3 4 -> 3
2 5 1 3 4 -> Too chaotic
5 1 2 3 7 8 6 4 -> Too chaotic
1 2 5 3 7 8 6 4 -> 7         	!
*/
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
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
