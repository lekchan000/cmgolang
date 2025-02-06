package main

import "fmt"

func main() {
	var pa1 product
	pa1.name = "Laptop"
	pa1.price = 5000
	pa1.stock = 50

	showProduct(pa1)
}

type product struct {
	name  string
	price float64
	stock int
}

func showProduct(pa1 product) {
	fmt.Println(pa1)
}
