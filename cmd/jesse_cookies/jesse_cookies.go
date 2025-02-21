package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc
https://www.youtube.com/watch?v=cJtq8eZCWpk

!!
https://medium.com/better-programming/using-heaps-to-speed-up-code-performance-in-go-b17ab07d2d5f
!!

*/

/*
 * Complete the 'cookies' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 */

func printList(mylist *list.List) {
	fmt.Println("printList-----")
	for element := mylist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	fmt.Println("-----printList")
}

func insertMixedCooky(cookie int, cookies *list.List) {
	for iteratedCookie := cookies.Front(); iteratedCookie != nil; iteratedCookie = iteratedCookie.Next() {
		iteratedCookieVal := iteratedCookie.Value.(int)
		if iteratedCookieVal >= cookie {
			cookies.InsertBefore(cookie, iteratedCookie)
			return
		}
	}
	cookies.PushBack(cookie)
}

func cookiesSlow(k int32, arr []int32) int32 {
	var iterations int32 = 0

	sweetEnough := func(k int32, cookies *list.List) bool {
		for cookie := cookies.Front(); cookie != nil; cookie = cookie.Next() {
			if cookie.Value.(int) < int(k) {
				return false
			}
		}

		return true
	}

	sortedCookies := make([]int, len(arr))
	for idx, item := range arr {
		sortedCookies[idx] = int(item)
	}
	sort.Ints(sortedCookies)

	cookieList := list.New()
	for _, cookie := range sortedCookies {
		cookieList.PushBack(cookie)
	}
	//printList(cookieList)

	if sweetEnough(k, cookieList) {
		return 0
	}

	found := false
	for {
		iterations++
		cookie1 := cookieList.Front()
		cookie1Val := cookie1.Value.(int)
		cookieList.Remove(cookie1)
		cookie2 := cookieList.Front()
		cookie2Val := cookie2.Value.(int)
		cookieList.Remove(cookie2)
		//printList(cookieList)

		mixedCookie := cookie1Val + 2*cookie2Val
		insertMixedCooky(mixedCookie, cookieList)
		//printList(cookieList)

		if sweetEnough(k, cookieList) {
			found = true
			break
		}

		if cookieList.Len() == 1 {
			break
		}
	}

	if found {
		return iterations
	} else {
		return -1
	}
}

// quicker that previous one but still not quick enough to pass all tests
func cookiesArrayBased(k int32, arr []int32) int32 {
	var iterations int32 = 0

	sweetEnough := func(k int32, cookies []int) bool {
		for _, cookie := range cookies {
			if cookie < int(k) {
				return false
			}
		}

		return true
	}

	sortedCookies := make([]int, len(arr))
	for idx, item := range arr {
		sortedCookies[idx] = int(item)
	}
	sort.Ints(sortedCookies)

	if sweetEnough(k, sortedCookies) {
		return 0
	}

	found := false
	for {
		iterations++
		twoLeftMostCookies := sortedCookies[0:2]
		sortedCookies = sortedCookies[2:]

		mixedCookie := twoLeftMostCookies[0] + 2*twoLeftMostCookies[1]
		sortedCookies = append(sortedCookies, mixedCookie)
		sort.Ints(sortedCookies)

		if sweetEnough(k, sortedCookies) {
			found = true
			break
		}

		if len(sortedCookies) == 1 {
			break
		}
	}

	if found {
		return iterations
	} else {
		return -1
	}
}

// see https://pkg.go.dev/container/heap
// https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc
// An IntMinHeap is a min-heap of ints.

type IntMinHeap []int32

func (h IntMinHeap) Len() int { return len(h) }

func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntMinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int32))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	lastVal := old[len(old)-1]
	*h = old[:len(old)-1]
	return lastVal
}

func (h IntMinHeap) AsSlice() []int32 {
	return []int32(h)
}

func cookies(k int32, arr []int32) int32 {
	if len(arr) == 0 {
		return -1
	}

	// initialize heap
	hv := IntMinHeap(arr)
	h := &hv
	heap.Init(h) // #complexity: O(n)

	var cont, min1, min2 int32
	for h.Len() > 1 {

		min1 = heap.Pop(h).(int32) // #complexity: O(log n)

		if min1 >= k {
			// if the lowest value is greater than k, we know that all remaining values will be too
			return cont
		}

		min2 = heap.Pop(h).(int32) // #complexity: O(log n)
		heap.Push(h, min1+2*min2)  // #complexity: O(log n)
		cont++

	}

	// total worst case complexity: O(n) + 3 O(log n) + 3 O(log n-2) + 3 O(log n-3)... = O(n)

	min1 = heap.Pop(h).(int32) // #complexity: O(1), since h has length = 1
	if min1 >= k {
		return cont
	}

	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var A []int32

	for i := 0; i < int(n); i++ {
		AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
		checkError(err)
		AItem := int32(AItemTemp)
		A = append(A, AItem)
	}

	result := cookies(k, A)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
