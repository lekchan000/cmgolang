package main

import "fmt"

func main() {
	fmt.Println("Begin")

	// Explicit Declaration
	var tmp1 int = 0
	var tmp2 string = "Hello"
	var tmp3 bool = true

	// const tmp4 int = 0

	// Implicit Declaration
	tmp5 := 0
	tmp6 := "codemobiles"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)

	fmt.Println(tmp5)
	fmt.Println(tmp6)
}
