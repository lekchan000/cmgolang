package main

import "fmt"

type product struct {
	name  string
	price float64
	stock int
}

func main() {
	var pa1 product
	pa1.name = "Laptop"
	pa1.price = 5000
	pa1.stock = 50

	showProduct(pa1)
	//updateProduct(&pa1)
	pa1.ClearProduct()
	showProduct(pa1)
}

func showProduct(pa1 product) {
	fmt.Println(pa1)
}

func (p product) ClearProduct() product {
	p.name = ""
	p.price = 0
	p.stock = 0
	return p
}

func updateProduct(pa1 *product) {
	pa1.price = 6000
	pa1.stock = 40
	fmt.Println(*pa1)
}
