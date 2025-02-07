package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer printfinish(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printfinish(i int) {
	fmt.Println("finish = ", i)
}
