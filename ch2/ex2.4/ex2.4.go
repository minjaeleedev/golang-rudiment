package main

import "fmt"

func main() {
	var a int = 3 // 기본 형태
	var b int     // 초깃값 생략. 초깃값은 타입별 기본값으로 대체
	var c = 4     // 타입 생략. 변수 타입은 우변 값의 타입이 됨
	d := 5        // 선언 대입문 := 사용. var 키워드와 타입 생략

	fmt.Println(a, b, c, d)

	var e = 3.1415     // e는 float64 타입으로 자동 지정
	f := 365           // f는 int 타입으로 자동 지정
	g := "hello world" // s는 string 타입으로 자동 지정

	fmt.Println(e, f, g)
}
