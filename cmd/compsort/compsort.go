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
 * Complete the 'countingSort' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func countingSort(arr []int32) []int32 {
	maxVal := arr[0]
	for _, value := range arr {
		if value > maxVal {
			maxVal = value
		}
	}

	//occArr := make([]int32, maxVal+1)
	occArr := make([]int32, 100)

	for _, val := range arr {
		occArr[val] += 1
	}
	/*
		result := make([]int32, 0)
		for idx, val := range occArr {
			for i := int32(0); i < val; i++ {
				result = append(result, int32(idx))
			}
		}
		return result
	*/
	return occArr
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := countingSort(arr)

	for i, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if i != len(result)-1 {
			fmt.Printf(" ")
		}
	}

	fmt.Printf("\n")
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
