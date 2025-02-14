package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
	gridRows := len(grid)
	gridCols := len(grid[0])

	// sort rows
	for idx, gridRow := range grid {
		r := strings.Split(gridRow, "")
		sort.Strings(r)
		grid[idx] = strings.Join(r, "")
	}

	// transponse matrix
	columns := make([]string, 0)
	for iCol := 0; iCol < gridCols; iCol++ {
		var col bytes.Buffer
		for iRow := 0; iRow < gridRows; iRow++ {
			col.WriteString(grid[iRow][iCol : iCol+1])
		}
		columns = append(columns, col.String())
	}

	// check sorting by column
	for _, column := range columns {
		r := strings.Split(column, "")
		sort.Strings(r)
		sortedColumns := strings.Join(r, "")
		if sortedColumns != column {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	grid1 := []string{"abc", "lmp", "qrt"}
	grid2 := []string{"mpxz", "abcd", "wlmf"}
	grid3 := []string{"abc", "hjk", "mpq", "rtv"}
	fmt.Println(gridChallenge(grid1))
	fmt.Println(gridChallenge(grid2))
	fmt.Println(gridChallenge(grid3))
}

func main2() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var grid []string

		for i := 0; i < int(n); i++ {
			gridItem := readLine(reader)
			grid = append(grid, gridItem)
		}

		result := gridChallenge(grid)

		fmt.Printf("%s\n", result)
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
