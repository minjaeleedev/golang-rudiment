package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300

	// len 함수 - 파라미터가 배열이면 배열요소 개수임
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
