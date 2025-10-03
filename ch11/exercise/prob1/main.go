package main

import "fmt"

type Product struct {
	Name        string
	Price       int
	ReviewScore float64
}

func main() {
	product := Product{"product", 21000, 4.8}
	fmt.Println(product)
}
