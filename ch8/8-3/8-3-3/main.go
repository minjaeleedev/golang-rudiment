package main

import "fmt"

func getMyAge() int {
	return 22
}

func printAgeV1() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("age is", age) // error - undefined: age
}

func printAgeV2() {
	// 초기문 + 비굣값 true 사용 - true 생략
	switch age := getMyAge(); {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}
}

func main() {
	printAgeV1()
	printAgeV2()
}
