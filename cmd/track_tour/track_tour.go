package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'truckTour' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY petrolpumps as parameter.
 */

func printList(mylist *list.List) {
	fmt.Println("-----")
	for element := mylist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	fmt.Println("-----")
}

func truckTour(petrolpumps [][]int32) int32 {
	var gasLevel int32 = 0
	pumpOffset := 0

	petrolPumpsList := list.New()
	for i := 0; i < len(petrolpumps); i++ {
		petrolPumpsList.PushBack(petrolpumps[i])
	}

	var found bool
	for i := 0; i < len(petrolpumps); i++ { // perform only as many iterations as we have pumps
		found = true
		//printList(petrolPumpsList)
		for element := petrolPumpsList.Front(); element != nil; element = element.Next() {
			elementVal, _ := element.Value.([]int32)
			gasTopUp := elementVal[0]
			distanceToNextPump := elementVal[1]
			gasLevel += gasTopUp
			gasLevel -= distanceToNextPump

			if gasLevel < 0 {
				found = false
				break
			}
		}

		if found {
			break
		} else {
			gasLevel = 0
			pumpOffset++
			// move first element to the end
			element := petrolPumpsList.Front()
			elementVal, _ := element.Value.([]int32)
			petrolPumpsList.Remove(element)
			petrolPumpsList.PushBack(elementVal)
		}
	}

	if found {
		return int32(pumpOffset)
	} else {
		return -1
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var petrolpumps [][]int32
	for i := 0; i < int(n); i++ {
		petrolpumpsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var petrolpumpsRow []int32
		for _, petrolpumpsRowItem := range petrolpumpsRowTemp {
			petrolpumpsItemTemp, err := strconv.ParseInt(petrolpumpsRowItem, 10, 64)
			checkError(err)
			petrolpumpsItem := int32(petrolpumpsItemTemp)
			petrolpumpsRow = append(petrolpumpsRow, petrolpumpsItem)
		}

		if len(petrolpumpsRow) != 2 {
			panic("Bad input")
		}

		petrolpumps = append(petrolpumps, petrolpumpsRow)
	}

	result := truckTour(petrolpumps)

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
