package main

import "fmt"

func nestedForLoop() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func triange() {
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func breakAndContinue() {
	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 {
				break
			}
		}
		b = 1
		dan++

		if dan == 10 {
			break
		}
	}

	fmt.Println("for문이 종료됐습니다.")
}

func main() {
	nestedForLoop()
	fmt.Println()
	triange()
	fmt.Println()
	breakAndContinue()
}
