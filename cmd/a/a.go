package main

import (
	"container/list"
	"fmt"
)

func printList(mylist *list.List) {
	fmt.Println("-----")
	for element := mylist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	fmt.Println("-----")
}
func main() {
	str := "abc"
	fmt.Println(str[1:2]) //2nd character
	fmt.Println(str[2:3]) //3nd character
	fmt.Println("Adam-----")
	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushBack(2)

	printList(mylist)

	element := mylist.Front()
	elementVal := element.Value
	mylist.Remove(element)
	mylist.PushBack(elementVal)

	printList(mylist)
}
