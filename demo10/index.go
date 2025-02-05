package main

import "fmt"

func printArray(arr [5]int) {
	for _, value := range arr {
		fmt.Println(value)
	}
}

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array elements:")
	printArray(arr1)
}
