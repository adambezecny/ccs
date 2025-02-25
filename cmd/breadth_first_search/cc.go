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
 * Complete the 'bfs' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n  - number of nodes
 *  2. INTEGER m  - number of edges
 *  3. 2D_INTEGER_ARRAY edges
 *  4. INTEGER s - starting node
 */

const notFound = "not_found" // default value

type Node struct {
	Value     int32
	Neighbors []int32
}

func pathNumber(path string) int32 {
	if path == notFound {
		return -1
	}
	return int32(strings.Count(path, "->") * 6)
}

func breadthFirstSearch(tree map[int32]Node, root, target int32) string {

	// check if root and target exist in the tree
	rootNode, rootExists := tree[root]
	_, targetExists := tree[target]
	if !rootExists || !targetExists {
		return notFound
	}

	// initialize the queue and push the root node
	q := list.New()
	q.PushBack(rootNode)

	// create a parent map to save the interactions and recreate the path
	parents := make(map[int32]int32) // initialize queue
	parents[root] = 0                // initialize root without any parents

	// while queue has elements, keep iterating
	for q.Len() > 0 {
		currentNode := q.Front().Value.(Node) // get first element
		q.Remove(q.Front())                   // remove first element from queue

		// compare if node is equals to target
		if currentNode.Value == target {
			// the target has been looked
			// reconstructing the path
			var route []int32
			for currentNode.Value > 0 {
				// recreating route
				route = append([]int32{currentNode.Value}, route...)
				// changing pointer
				currentNode.Value = parents[currentNode.Value]
			}

			// returning path result
			var stringSlice []string
			for _, num := range route {
				stringSlice = append(stringSlice, strconv.Itoa(int(num)))
			}

			return strings.Join(stringSlice, "->")
		}

		// iterate neighbors
		for _, neighbor := range currentNode.Neighbors {
			// check if the neighbor has not already been visited
			if _, visited := parents[neighbor]; !visited {
				parents[neighbor] = currentNode.Value // add neighbor to parents map associated to current node value
				q.PushBack(tree[neighbor])            // add neighbor to the end of the queue
			}
		}
	}

	return notFound
}

func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	tree := make(map[int32]Node)

	for i := int32(1); i <= n; i++ {
		tree[i] = Node{Value: i, Neighbors: []int32{}}
	}

	for _, edge := range edges {
		nodeFrom := edge[0]
		nodeTo := edge[1]

		neighbours := tree[nodeFrom].Neighbors
		neighbours = append(neighbours, nodeTo)

		tree[nodeFrom] = Node{Value: nodeFrom, Neighbors: neighbours}
	}

	result := make([]int32, 0)

	for i := int32(1); i <= n; i++ {
		if i == s {
			continue
		}

		path := breadthFirstSearch(tree, s, i)
		result = append(result, pathNumber(path))
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		var edges [][]int32
		for i := 0; i < int(m); i++ {
			edgesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var edgesRow []int32
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int32(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != 2 {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := bfs(n, m, edges, s)

		for i, resultItem := range result {
			fmt.Printf("%d", resultItem)

			if i != len(result)-1 {
				fmt.Printf(" ")
			}
		}

		fmt.Printf("\n")
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
