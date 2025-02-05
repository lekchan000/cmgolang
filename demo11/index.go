package main

import "fmt"

func main() {
	// showSlice(make([]int, 3, 5))
	var numbers []int
	numbers = append(numbers, 555, 666, 777)
	showSlice(numbers)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
