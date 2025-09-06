package main

import "fmt"

func main() {
	var a = 123
	var b int = 4567
	f := 3.14159269

	// fmt.Printf와 서식문자를 이용해서 원하는 출력하기
	fmt.Printf("%6d\n", a)
	fmt.Printf("%06d\n", b)
	fmt.Printf("%6.2f\n", f) // 3.14. 소수점 이하 2자리 표시
}
