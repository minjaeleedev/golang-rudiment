package main

import "fmt"

func main() {
	var x int8 = 4
	var y int8 = 64

	fmt.Println("left shift")
	fmt.Printf("x:%08b x<<2: %08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2: %08b y<<2: %d\n", y, y<<2, y<<2)

	fmt.Println("right shift")
	var x2 int8 = 16
	var y2 int8 = -128
	var z2 int8 = -1   // 부호 있는 정수, 최상위 비트 1
	var w2 uint8 = 128 // 부호 없는 정수, 최상위 비트 1
	var v2 uint8 = 64  // 부호 없는 정수, 최상위 비트 0
	fmt.Printf("x2:%08b x2>>2: %08b x2>>2: %d\n", x2, x2>>2, x2>>2)
	fmt.Printf("y2:%08b y2>>2: %08b y2>>2: %d\n", uint8(y2), uint8(y2>>2), y2>>2)
	fmt.Printf("z2:%08b z2>>2: %08b z2>>2: %d\n", uint8(z2), uint8(z2>>2), z2>>2)
	fmt.Printf("w2:%08b w2>>2: %08b w2>>2: %d\n", uint8(w2), uint8(w2>>2), w2>>2)
	fmt.Printf("v2:%08b v2>>2: %08b v2>>2: %d\n", uint8(v2), uint8(v2>>2), v2>>2)
}
