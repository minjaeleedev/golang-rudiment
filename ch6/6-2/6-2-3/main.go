package main

import "fmt"

const (
	Red   int = iota
	Blue  int = iota
	Green int = iota
)

const (
	// 첫 번째 값과 똑같은 규칙이 적용되면 타입과 iota 키워드 생략 가능
	C1 uint = iota + 1
	C2      // C2 uint = iota + 1 이고, iota가 1로 증가했으므로 2
	C3
)

const (
	BitFlag1 uint = 1 << iota // 1 << 0 이므로 1
	BitFlag2                  // 1 << 1
	BitFlag3                  // 1 << 2
	BitFlag4
)

const (
	A int = iota // 0
	B            // 1
)

func main() {
	fmt.Println(Red)
	fmt.Println(Blue)
	fmt.Println(Green)

	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(C3)

	fmt.Println(BitFlag1)
	fmt.Println(BitFlag2)
	fmt.Println(BitFlag3)
	fmt.Println(BitFlag4)

	fmt.Println(A)
	fmt.Println(B)
}
