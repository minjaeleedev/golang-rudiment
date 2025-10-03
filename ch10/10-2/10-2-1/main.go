package main

import "fmt"

func main() {
	var nums [5]int
	for i := 0; i < 5; i++ {
		fmt.Println(nums[i])
	}
	fmt.Println()

	days := [3]string{"monday", "tuesday", "wednesday"}
	for i := 0; i < 3; i++ {
		fmt.Println(days[i])
	}
	fmt.Println()

	var temps [5]float64 = [5]float64{24.3, 26.7}
	for i := 0; i < 5; i++ {
		fmt.Println(temps[i])
	}
	fmt.Println()

	var s = [5]int{1: 10, 3: 30}
	for i := 0; i < 5; i++ {
		fmt.Println(s[i])
	}
	fmt.Println()

	var x = [...]int{10, 20, 30}
	for i := 0; i < 3; i++ {
		fmt.Println(x[i])
	}
}
