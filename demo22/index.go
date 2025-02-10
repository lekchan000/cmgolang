package main

import (
	"fmt"
	"time"
)

func task(c, quit chan string) {
	for {
		select {
		case c <- "running":
		case <-quit:
			fmt.Println("quit")
			return
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan string)
	quit := make(chan string)
	go func() { // waiting for result from task
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// quit <- "done"
	}()

	task(c, quit)
}
