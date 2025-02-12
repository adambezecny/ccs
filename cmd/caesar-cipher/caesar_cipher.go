package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	letters := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z",
	}

	var indexOf = func(element string) int {
		for i, v := range letters {
			if v == element {
				return i
			}
		}
		return -1
	}

	var shiftBy = func(element string, shiftBy int) string {
		position := indexOf(element)
		if position == -1 {
			return element
		}

		// 0+3
		if position+shiftBy < len(letters) {
			return letters[position+shiftBy]
		} else {
			moveToEndCount := len(letters) - 1 - position
			moveFromBeginningCount := shiftBy - moveToEndCount - 1
			return letters[moveFromBeginningCount]
		}
	}

	result := ""
	for _, char := range s {
		isUpper := unicode.IsUpper(char)
		if !isUpper {
			result += shiftBy(fmt.Sprintf("%c", char), int(k))
		} else {
			result += strings.ToUpper(
				shiftBy(
					strings.ToLower(fmt.Sprintf("%c", char)), int(k),
				),
			)
		}

	}
	return result
}

func noop(_ int32) {
	//do nothing
}

/*func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

  	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

	noop(n)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}*/

/*
159357lcfd
98
->
159357fwzx
*/
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	noop(n)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Printf("%s\n", result)
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
