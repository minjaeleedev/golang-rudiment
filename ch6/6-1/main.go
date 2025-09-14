package main

import "fmt"

func main() {
	const C int = 10 // 상수 선언
	fmt.Println(C)

	var b int = C * 20
	fmt.Println(b)
	// C = 20  cannot assign to C (neither addressable nor a map index expression)
	// fmt.Println(&C) invalid operation: cannot take address of C (constant 10 of type int)
}
