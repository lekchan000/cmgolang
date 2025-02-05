package main

import "fmt"

func main() {
	coruse := []string{"Python", "Java", "Go", "C++", "C#", "JavaScript", "Ruby", "PHP", "Swift", "Kotlin"}
	for index, item := range coruse {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	for _, item := range coruse {
		fmt.Printf("%s\n", item)
	}
}
