package main

import "fmt"

func main() {
	str := "adxm"
	fmt.Println(str[:0] + str[1:])
	fmt.Println(str[:1] + str[2:])
	fmt.Println(str[:2] + str[3:])
	fmt.Println(str[:3] + str[4:])

}
