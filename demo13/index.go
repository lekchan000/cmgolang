package main

import "fmt"

func main() {
	var num = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(num["one"], num["two"], num["three"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println(colors["red"], colors["green"], colors["blue"])
}
