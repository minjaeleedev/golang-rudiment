package main

import "fmt"

func breakWithFlag() {
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

func breakWithLabel() {
	a := 1
	b := 1

OuterFor: // 레이블 정의
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
}

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}

	return 0, false
}

func refactorWithFunction() {
	a := 1
	b := 0
	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found {
			break
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

func main() {
	breakWithFlag()
	breakWithLabel()
	refactorWithFunction()
}
