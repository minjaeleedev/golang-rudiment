package main

import "fmt"

// 상숫값에 코드를 부여한다
const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	switch animal {
	case Pig:
		fmt.Println("꿀꿀")
	case Cow:
		fmt.Println("음메")
	case Chicken:
		fmt.Println("꼬끼오")
	default:
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}
