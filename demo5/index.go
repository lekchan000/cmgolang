package main

import "fmt"

func main() {
	fn1()
	fn2("Good Morning My neighbor")
	fn3("Good Morning My neighbor", 2)
}

func fn1() {
	fmt.Println("CodeMobiles")
}
func fn2(msg string) {
	fmt.Println("msg : ", msg)
}
func fn3(title string, version int) {
	fmt.Println(title, "\n", version)
}
