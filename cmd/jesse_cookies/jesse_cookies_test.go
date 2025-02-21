package main

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestCookiesSlow(t *testing.T) {
	assert.Equal(t, int32(4), cookiesSlow(9, []int32{2, 7, 3, 6, 4, 6}))
	assert.Equal(t, int32(2), cookiesSlow(7, []int32{1, 2, 3, 9, 10, 12}))
	assert.Equal(t, int32(-1), cookiesSlow(1000, []int32{1, 2, 3}))
	assert.Equal(t, int32(0), cookiesSlow(1, []int32{2, 3, 4}))
}

func TestCookiesArrayBased(t *testing.T) {
	assert.Equal(t, int32(4), cookiesArrayBased(9, []int32{2, 7, 3, 6, 4, 6}))
	assert.Equal(t, int32(2), cookiesArrayBased(7, []int32{1, 2, 3, 9, 10, 12}))
	assert.Equal(t, int32(-1), cookiesArrayBased(1000, []int32{1, 2, 3}))
	assert.Equal(t, int32(0), cookiesArrayBased(1, []int32{2, 3, 4}))
}

func TestCookies(t *testing.T) {
	assert.Equal(t, int32(4), cookies(9, []int32{2, 7, 3, 6, 4, 6}))
	assert.Equal(t, int32(2), cookies(7, []int32{1, 2, 3, 9, 10, 12}))
	assert.Equal(t, int32(-1), cookies(1000, []int32{1, 2, 3}))
	assert.Equal(t, int32(0), cookies(1, []int32{2, 3, 4}))
}

func BenchmarkCookies(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	list := make([]int32, 0)
	for i := 0; i < 10_000; i++ {
		randNum := rand.Intn(201-100) + 100
		list = append(list, int32(randNum))
	}

	for i := 0; i < b.N; i++ {
		cookies(125, list)
	}
}

func BenchmarkCookiesSlow(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	list := make([]int32, 0)
	for i := 0; i < 10_000; i++ {
		randNum := rand.Intn(201-100) + 100
		list = append(list, int32(randNum))
	}

	for i := 0; i < b.N; i++ {
		cookiesSlow(125, list)
	}
}

func TestInsertMixedCooky(t *testing.T) {
	cookieList := list.New()
	cookieList.PushBack(4)
	cookieList.PushBack(6)
	cookieList.PushBack(6)
	cookieList.PushBack(7)

	cookieListToArr := func(cookies *list.List) []int {
		var cookieArr []int
		for element := cookies.Front(); element != nil; element = element.Next() {
			cookieArr = append(cookieArr, element.Value.(int))
		}
		return cookieArr
	}

	insertMixedCooky(8, cookieList)
	assert.Equal(t, []int{4, 6, 6, 7, 8}, cookieListToArr(cookieList))

	insertMixedCooky(6, cookieList)
	assert.Equal(t, []int{4, 6, 6, 6, 7, 8}, cookieListToArr(cookieList))

	insertMixedCooky(5, cookieList)
	assert.Equal(t, []int{4, 5, 6, 6, 6, 7, 8}, cookieListToArr(cookieList))
}
