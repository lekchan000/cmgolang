package main

import "fmt"

func main() {
	fn1()
	fn2("Good Morning My neighbor")
	fn3("Good Morning My neighbor", 2)

	const a, b = 2, 3
	fmt.Printf("%d+%d=%d", a, b, sum(a, b))
	fmt.Printf("\n")

	var x, y int = swap(1, 0)
	fmt.Println(x, y)
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

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}
