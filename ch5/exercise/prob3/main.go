package prob3

import "fmt"

func F(n int) int {
	// 탈출 조건
	// if n == 0 {
	// 	return 0
	// }

	// if n == 1 {
	// 	return 1
	// }
	if n < 2 {
		return n
	}

	return F(n-2) + F(n-1)
}

func main() {
	fmt.Println(F(9))
}
