package main

import "fmt"

var arr [5]int

func main() {
	arr[0] = 10
	arr[1] = 15

	fmt.Println(arr[0], arr[1], arr)
}
