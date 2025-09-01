package main

import "fmt"

func main() {
	// 에러 코드 - 주석풀면 에러 확인 가능
	// x := 3              // int
	// var y float64 = 3.5 // float64

	// var z int = y
	// w := x * y

	// var v int64 = 7
	// u := x * v

	// var t int = y * 3
	// fmt.Println(t, u, v, w, x, y, z)

	a := 3
	var b float64 = 3.5

	var c int = int(b)  // float64에서 int로 변환
	d := float64(a * c) // int * int = int. int에서 float64로 변환
	fmt.Println(d)

	var e int64 = 7
	f := int64(d) * e // float64 to int64

	var g int = int(b * 3)
	var h int = int(b) * 3

	fmt.Println(g, h, f) // answer => 10 9 63
}
