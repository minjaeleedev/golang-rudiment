package main

import "fmt"

func main() {
	fmt.Println(3*4^7<<2+3*5 == 7)
	// 12^7<<2+3*5 == 7
	// 12^28+3*5 == 7
	// 12^28+15 == 7
	// 16+15 == 7
	// 31 == 7
	// false
	// for readability
	fmt.Println((3*4 ^ (7 << 2) + (3 * 5)) == 7)
}
