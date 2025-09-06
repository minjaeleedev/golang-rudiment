package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)
	// n : 성공적으로 입력한 값 개수
	// err : 입력시 발생한 에러
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
