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
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	sums := [5]int64{0, 0, 0, 0, 0}
	var minVal int64 = 0
	var maxVal int64 = 0

	for i := 0; i <= 4; i++ {
		for idx, item := range arr {
			if idx != i {
				sums[i] += int64(item)
			}
		}
	}

	// find min
	minVal = int64(sums[0])
	for _, value := range sums {
		if value < minVal {
			minVal = value
		}
	}

	// find max
	maxVal = int64(sums[0])
	for _, value := range sums {
		if value > maxVal {
			maxVal = value
		}
	}

	//fmt.Printf("%v", sums)
	fmt.Printf("%d %d", minVal, maxVal)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
