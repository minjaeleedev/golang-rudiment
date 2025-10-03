package main

import "fmt"

func main() {
	var p *int
	var a int
	fmt.Scanln(&a)
	if a%2 == 0 {
		p = &a
	}

	if p != nil {
		fmt.Printf("p: %p\n", p)
	} else {
		fmt.Printf("p is nil: %p\n", p)
	}
}
