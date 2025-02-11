package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

	plusMinus(arr)
}

func plusMinus(arr []int32) {
	positives := make([]int32, 0)
	negatives := make([]int32, 0)
	zeros := make([]int32, 0)

	for _, item := range arr {
		if item > 0 {
			positives = append(positives, item)
		} else if item < 0 {
			negatives = append(negatives, item)
		} else {
			zeros = append(zeros, item)
		}
	}

	fmt.Printf("%.6f\n", float32(len(positives))/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(len(negatives))/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(len(zeros))/float32(len(arr)))
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
