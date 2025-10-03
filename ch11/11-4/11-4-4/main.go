// memory padding
package main

import (
	"fmt"
	"unsafe"
)

type UserV1 struct {
	A int8
	B int
	C int8
	D int
	E int8
}

// 패딩 규칙 고려한 필드 배치
type UserV2 struct {
	A int8
	C int8
	E int8
	B int
	D int
}

func main() {
	user := UserV1{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user)) // 40

	user2 := UserV2{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user2)) // 24
}
