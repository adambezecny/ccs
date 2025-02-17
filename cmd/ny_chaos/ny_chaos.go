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

func minimumBribesNotWorking(q []int32) {
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

func minimumBribes(q []int32) {
	var smallerThanCount = func(nums []int32, position int) int {
		counter := 0
		for i := len(nums) - 1; i > position; i-- {
			if nums[i] < nums[position] {
				counter++
			}
		}
		return counter
	}

	// first find all numbers not on their positions
	outOfPosition := make([]int, 0)
	for i := 0; i < len(q); i++ {
		if int(q[i]) != i+1 {
			outOfPosition = append(outOfPosition, i)
		}
	}

	// then for each out of position number find how many smaller numbers are to the left
	// i.e. how many smaller numbers this out of pos number has overtaken
	counter := 0
	for _, val := range outOfPosition {
		inc := smallerThanCount(q, val)
		if inc > 2 {
			fmt.Println("Too chaotic")
			return
		}
		counter += inc
	}
	fmt.Println(counter)
}

func bubbleSort(numbers []int) {
	counter := 0
	length := len(numbers)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
				counter++
			}
		}
	}
	fmt.Printf("number of swPS: %d\n", counter)
}

/*
2 1 5 3 4 -> 3
2 5 1 3 4 -> Too chaotic
5 1 2 3 7 8 6 4 -> Too chaotic

1 2 5 3 7 8 6 4 -> 7         	!minimumBribesNotWorking returns 6

									(does not count it since it ended up on worse position than 6 actually,
	 									but it still bribed and did the swap, see bellow!)

1 2 3 4 5 6 7 8
1 2 3 5 4 6 7 8  +1(5)
1 2 5 3 4 6 7 8  +1(5)
1 2 5 3 6 4 7 8  +1(6)
1 2 5 3 6 7 4 8  +1(7)
1 2 5 3 6 7 8 4  +1(8)
1 2 5 3 7 6 8 4  +1(7)
1 2 5 3 7 8 6 4  +1(8) -> 7 swaps

number of swaps for any given number = number of smaller numbers on higher positions (to the right)
e.g. 6 has one smaller number right to itself, 4 -> one swap, 5 has two such numbers (3,4)-> two swaps
*/
func main() {
	/*arr := [8]int{1, 2, 5, 3, 7, 8, 6, 4}
	bubbleSort(arr[:])
	return*/
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
