package main

import "fmt"

func main() {
	// "Hello World", 0, 30 : Literal.
	var str string = "Hello World"
	var i int = 0
	i = 30

	fmt.Println(str)
	fmt.Println(i)

	// Go에서 상수는 리터럴과 같이 취급됨
	const PI = 3.14
	var a int = PI * 100 // 컴파일 타임에 실제 결괏값 리터럴(314)로 변환
	fmt.Println(a)
	// 상수의 메모리 주솟값에 접근할 수 없는 이유 : 컴파일 타임에 리터럴로 전환
	// 실행 파일에 값 형태로 쓰임
	// 따라서 상수는 동적 할당 메모리 영역을 사용하지 않음
}
