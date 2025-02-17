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

func superDigitRecursive(n string, k int32) int32 {
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

	return superDigitRecursive(strconv.Itoa(sum), 1)
}

func superDigit(n string, k int32) int32 {
	str := ""

	var sumDigits = func(digits []string) int {
		sum := 0
		for _, digitStr := range digits {
			digit, err := strconv.Atoi(digitStr)
			if err != nil {
				panic(err)
			}

			sum += digit
		}
		return sum
	}

	var isSingleDigit = func(digit int) bool {
		return len(strconv.Itoa(digit)) == 1
	}

	/*for i := 0; i < int(k); i++ {
		str += n
	}*/
	str = n

	sumOfDigits := 0
	digits := strings.Split(str, "")
	for {
		sumOfDigits = sumDigits(digits)
		if k > 1 {
			sumOfDigits *= int(k)
			k = 1
		}
		if isSingleDigit(sumOfDigits) {
			break
		} else {
			digits = strings.Split(strconv.Itoa(sumOfDigits), "")
		}
	}

	return int32(sumOfDigits)
}

/*
3546630947312051453014172159647935984478824945973141333062252613718025688716704470547449723886626736 100000
expected output is 5

148 3 -> expected output is 3
9875 4 -> 8
123 3 -> 9
*/
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
