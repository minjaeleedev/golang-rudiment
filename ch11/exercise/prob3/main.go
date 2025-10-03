package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A int8
	B int
	C float64
	D uint16
	E int
	F float32
	G int8
}

type PaddingV2 struct {
	A int8    // 1
	G int8    // 1
	D uint16  // 2
	F float32 // 4
	B int     // 8
	E int     // 8
	C float64 // 8
}

func main() {
	padding := Padding{
		1, 2, 3, 4, 5, 6, 7,
	}

	paddingV2 := PaddingV2{
		1, 2, 3, 4, 5, 6, 7,
	}

	fmt.Println(unsafe.Sizeof(padding))
	fmt.Println(unsafe.Sizeof(paddingV2))
}
