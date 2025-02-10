package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1           //sent
	fmt.Println(<-ch) //received

	ch <- 2           //sent
	fmt.Println(<-ch) //received

	time.Sleep(1 * time.Second)
}
