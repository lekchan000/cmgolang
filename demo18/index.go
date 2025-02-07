package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func myinitSeq() func() string {
	j := 0
	return func() string {
		j++
		return fmt.Sprintf("Y = %d ", j)
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
	fmt.Println(newInts())

	newInts2 := intSeq()
	fmt.Println(newInts2())
	fmt.Println(newInts2())
	fmt.Println(newInts2())

	newString := myinitSeq()
	fmt.Println(newString())
	fmt.Println(newString())
	fmt.Println(newString())
}
