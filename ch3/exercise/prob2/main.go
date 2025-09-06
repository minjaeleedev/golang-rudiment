package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanln(a, b)
	fmt.Println(a, b)

	// Scanln의 인자는 변수의 주소를 받아야 한다.
}
