package main

import "fmt"

func main() {
	var a int8 = 30

	a <<= 2 // 120
	a += 8  // -128
	fmt.Println(a)
}
