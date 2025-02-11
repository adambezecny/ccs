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
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

// https://www.youtube.com/watch?v=1jGAutnHuYM (see min 9:46)
// https://www.youtube.com/watch?v=4rin1enhuQQ
func flippingMatrix(matrix [][]int32) int32 {
	type mirrorMember struct {
		X int
		Y int
	}

	type mirror struct {
		M1 mirrorMember
		M2 mirrorMember
		M3 mirrorMember
		M4 mirrorMember
	}

	mirrors := make([]mirror, 0)

	matrixSideLen := len(matrix[0])

	//no of mirrors:
	/*
		MATRIX 4x4 -> 2x2=4 mirrors
		matrix 6,6 -> 3x3=9 mirrors
	*/
	//noOfMirrors := matrixSideLen / 2 * matrixSideLen / 2

	// first determine all mirrors

	for x := 0; x < matrixSideLen/2; x++ {
		for y := 0; y < matrixSideLen/2; y++ {
			mirrors = append(mirrors, mirror{
				M1: mirrorMember{X: x, Y: y},                                         // 0,0
				M2: mirrorMember{X: x, Y: matrixSideLen - y - 1},                     // 0,5
				M3: mirrorMember{X: matrixSideLen - x - 1, Y: y},                     // 5,0
				M4: mirrorMember{X: matrixSideLen - x - 1, Y: matrixSideLen - y - 1}, // 5,5
			})
		}
	}

	//then determine biggest number of every mirror
	mirrorsMaxVals := make([]int32, 0)
	for _, m := range mirrors {
		mirrorValues := [4]int32{
			matrix[m.M1.X][m.M1.Y],
			matrix[m.M2.X][m.M2.Y],
			matrix[m.M3.X][m.M3.Y],
			matrix[m.M4.X][m.M4.Y],
		}

		maxVal := mirrorValues[0]
		for _, value := range mirrorValues {
			if value > maxVal {
				maxVal = value
			}
		}

		mirrorsMaxVals = append(mirrorsMaxVals, maxVal)
	}

	// summarize max vals for each mirror
	var sum int32 = 0
	for _, val := range mirrorsMaxVals {
		sum += val
	}

	return sum
}

// 1 2 then 4x4 matrix
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var matrix [][]int32
		for i := 0; i < 2*int(n); i++ {
			matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var matrixRow []int32
			for _, matrixRowItem := range matrixRowTemp {
				matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
				checkError(err)
				matrixItem := int32(matrixItemTemp)
				matrixRow = append(matrixRow, matrixItem)
			}

			if len(matrixRow) != 2*int(n) {
				panic("Bad input")
			}

			matrix = append(matrix, matrixRow)
		}

		result := flippingMatrix(matrix)

		fmt.Printf("%d\n", result)
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
