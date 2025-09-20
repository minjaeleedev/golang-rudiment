package main

import "fmt"

// break
func switchWithBreak() {
	fmt.Println("break example")
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		// break // redundant break statement warning.
	case 2:
		fmt.Println("a == 2")
		// break
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}

func switchWithFallThrough() {
	fmt.Println("fallthough example")
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		// break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}

func main() {
	switchWithBreak()
	switchWithFallThrough()
}
