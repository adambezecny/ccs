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
 * Complete the 'towerBreakers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 */

// https://www.youtube.com/watch?v=jOxTTE3IkjE
/*
n no of towers
m tower height

if m=1 player 1 always loses
if n=1 player 2 always loses (because player 1 removes all but 1st element leaving single element which cannot be removed)

n is even player 1 always loses
n is odd player 2 always loses

func returns the winner of the game
*/
func towerBreakers(n int32, m int32) int32 {
	if m == 1 || n%2 == 0 {
		return 2
	} else {
		return 1
	}
}

func towerBreakersNotWorking(n int32, m int32) int32 {
	towers := make([]int32, n)

	for i := 0; i < int(n); i++ {
		towers[i] = m
	}

	playerOneTurn := true

	playerOneDidMove := false
	playerTwoDidMove := false

	var setCurrentPlayerDidMove = func() {
		if playerOneTurn {
			playerOneDidMove = true
			playerTwoDidMove = false
		} else {
			playerOneDidMove = false
			playerTwoDidMove = true
		}
	}

	for {
		for idx, tower := range towers {
			if tower%2 == 0 {
				towers[idx] /= 2
				setCurrentPlayerDidMove()
				break
			} else if tower != 1 {
				towers[idx] = 1
				setCurrentPlayerDidMove()
				break
			}
		}

		if playerOneTurn && !playerOneDidMove {
			return 2
		}

		if !playerOneTurn && !playerTwoDidMove {
			return 1
		}

		playerOneTurn = !playerOneTurn
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		result := towerBreakers(n, m)

		fmt.Printf("%d\n", result)
	}
}

/*func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		result := towerBreakers(n, m)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}*/

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
