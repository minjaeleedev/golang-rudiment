package main

import "fmt"

func main() {
	var x int8 = 1                 // 0000 0001
	x <<= 7                        // 1000 0000
	x >>= 7                        // 1111 1111 오른쪽으로 shift할 때마다 1로 채워진다
	fmt.Printf("%08b\n", uint8(x)) // 1111 1111
	fmt.Printf("%d\n", x)          // -1
}
