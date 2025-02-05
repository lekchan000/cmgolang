package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
}

func fnFor() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Index %d\n", i)
	}
}

func fnWhile() {
	i := 0
	for i <= 10 {
		fmt.Printf("Index %d\n", i)
		i++
	}
}
