package main

import "fmt"

func main() {
	var f32 float32 = 16777216 // 2^24
	fmt.Printf("float32: %.0f + 1 = %.0f\n", f32, f32+1)

	var f64 float64 = 16777216
	fmt.Printf("float64: %.0f + 1 = %.0f\n", f64, f64+1)

	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) // 4.266663e+06, 실제 값 4266663.334329 (7자리 제한)
	fmt.Println(d) // 1.2799989e+07, 실제 값 12799990.002987
}
