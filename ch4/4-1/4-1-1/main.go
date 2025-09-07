package main

import "fmt"

func main() {
	var x int32 = 7
	var y int32 = 3

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("s * t = ", s*t) // 15.700001
	fmt.Println("s / t = ", s/t)
}
