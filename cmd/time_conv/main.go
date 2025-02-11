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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	isAm := strings.HasSuffix(s, "AM")
	splitted := strings.Split(s, ":")

	hourSegment := splitted[0]
	minSegment := splitted[1]
	secSegment := splitted[2]

	convertedHour := hourSegment
	secSegmentWithoutAmPm := secSegment[:len(secSegment)-2]

	if isAm /* is AM time */ {
		if hourSegment == "12" {
			convertedHour = "00"
		}
	} else /* is PM time */ {
		if hourSegment == "12" {
			convertedHour = "12"
		} else {
			if i, err := strconv.Atoi(hourSegment); err != nil {
				panic(err)
			} else {
				convertedHour = strconv.Itoa(i + 12)
			}
		}
	}
	return fmt.Sprintf(
		"%s:%s:%s",
		convertedHour,
		minSegment,
		secSegmentWithoutAmPm,
	)
}

/*func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}*/

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	s := readLine(reader)
	result := timeConversion(s)
	fmt.Print(result)
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
