package main

import (
	"fmt"

	"github.com/lekchan000/cmgolang/demo15/employee"
)

func main() {
	e := employee.New(
		"John", "Doe", 10, 5)
	fmt.Println(e.LeavesRemaining())
}
