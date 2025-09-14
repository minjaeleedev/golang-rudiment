package main

import "fmt"

const PI = 3.14              // untyped constant
const PI2 = 3.14159          // untyped constant. more precision
const FloatPI float64 = 3.14 // float64 type constant

func main() {
	var a int = PI * 100
	// var b int = FloatPI * 100  타입 오류 발생
	// var c int = PI2 * 100 타입 오류 발생

	fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
}
