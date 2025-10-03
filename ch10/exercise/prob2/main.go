package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr [3][2][5]float64

	fmt.Println(unsafe.Sizeof(arr)) // 8*30 = 240
}
