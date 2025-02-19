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
https://tomoharutsutsumi.medium.com/how-to-avoid-o-n%C2%B2-60eaa61f523a
https://dev.to/leandronsp/how-to-reduce-the-time-complexity-of-nested-loops-1lkd

two pointers in loop approach: https://loopccew.medium.com/two-pointer-approach-5de851731ce8
*/

/*
 * Complete the 'pairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY arr
 */
// https://www.youtube.com/watch?v=xtOBbWMc0ZA
func pairs(k int32, arr []int32) int32 {
	arrSorted := make([]int, len(arr))
	for idx, item := range arr {
		arrSorted[idx] = int(item)
	}
	sort.Ints(arrSorted)

	counter := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arrSorted[j]-arrSorted[i] == int(k) {
				counter++
			}
			if arrSorted[j]-arrSorted[i] > int(k) {
				break
			}
		}
	}

	return int32(counter)
}

// very slow, O(n^2) complexity
func pairsSlow(k int32, arr []int32) int32 {
	arrLen := len(arr)
	counter := 0
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen; j++ {
			if i == j {
				continue
			}
			if arr[i]-arr[j] == k {
				counter++
			}
		}
	}

	return int32(counter)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)
	fmt.Printf("%d\n", result)
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
