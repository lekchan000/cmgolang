package main

import "fmt"

func main() {
	msg := "Hello, World!"
	var msgPtr *string = &msg
	fmt.Println("Value of msg: ", msg)
	fmt.Println("Value of msgPtr: ", msgPtr)
	fmt.Println("Value of *msgPtr: ", *msgPtr)

	changeMessage(&msg) // change the value of msg
	fmt.Println("Value of msg: ", msg)
}

func changeMessage(aPointer *string) {
	*aPointer = "Changed!"

}
