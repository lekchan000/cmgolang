package main

import "fmt"

func main() {
	someValue := 11
	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue > 10 && someValue < 20 {
		fmt.Println("someValue > 10 && someValue < 20")
	} else {
		fmt.Println("Not someValue > 10 && someValue < 20")
	}

	if rslt := getResult(); rslt == 10 {
		fmt.Println("rslt == 10")
	} else {
		fmt.Println("rslt != 10")
	}
	switchCase()
}

// if else shortform
func getResult() int {
	return 5
}

func switchCase() {
	index := 2
	switch index {

	case 0:
		fmt.Println("Zero")

	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")

	default:
		fmt.Println("Default case")
	}
}
