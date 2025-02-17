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
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func superDigit(n string, k int32) int32 {
	str := ""

	for i := 0; i < int(k); i++ {
		str += n
	}

	digits := strings.Split(str, "")

	if len(digits) == 1 {
		digit, err := strconv.Atoi(digits[0])
		if err != nil {
			panic(err)
		}
		return int32(digit)
	}

	sum := 0
	for _, digitStr := range digits {
		digit, err := strconv.Atoi(digitStr)
		if err != nil {
			panic(err)
		}

		sum += digit
	}

	return superDigit(strconv.Itoa(sum), 1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

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
