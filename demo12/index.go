package main

import "fmt"

func main() {
	var numbers []int
	numbers = append(numbers, 555, 666, 777)
	showSlice(numbers)

	remove(numbers, 0)
	remove(numbers, 0)
	remove(numbers, 0)
	showSlice(numbers)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
