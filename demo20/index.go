package main

import (
	"fmt"
	"time"
)

func routine1(c chan int, payload int) {
	c <- payload         //send payload to channel
	fmt.Println("step1") //print sent message

	c <- payload         //send payload to channel
	fmt.Println("step2") //print sent message

	c <- payload         //send payload to channel
	fmt.Println("step3") //print sent message
}

func main() {
	ch := make(chan int, 5) //create a channel with buffer size 3

	go routine1(ch, 1) //send

	fmt.Println(<-ch) //received
	fmt.Println(<-ch) //received
	fmt.Println(<-ch) //received

	time.Sleep(1 * time.Second) //wait for 1 second
}
