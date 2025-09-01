package main

import "fmt"

func main() {
	var a int32 = 360
	var b int8 = int8(a)
	fmt.Println(b)
	// a : 00000000 00000000 00000001 01101000 = 360
	// b : 00000000 00000000 00000000 01101000 = 104
}
