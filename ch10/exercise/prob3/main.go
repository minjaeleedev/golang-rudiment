package main

import "fmt"

func ChangeArray(arr [5]int) {
	arr[3] = 3000
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	ChangeArray(a)

	fmt.Println(a[3]) // 4.
	// 함수 인수로 값이 복사된다.
	// 변수 a와 ChangeArray의 변수 arr는 서로 다른 메모리 주소를 가진 다른 배열
	// arr[3] 바꿔도 a[3]은 바뀌지 않음
}
