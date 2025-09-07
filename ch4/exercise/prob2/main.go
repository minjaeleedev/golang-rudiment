package main

import "fmt"

func main() {
	var a uint8
	a |= 2
	a |= 4
	a |= 8 // 0000 1110

	var b uint8
	b = 4 // 0000 0100

	a &^= b        // 0000 1010
	fmt.Println(a) // 10
}
