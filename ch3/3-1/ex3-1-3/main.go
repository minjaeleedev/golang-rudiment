package main

import "fmt"

func main() {
	var a = 123
	var b = 456

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b)
	fmt.Printf("%-5d, %-05d\n", a, b)

	var c = 123456789
	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)

	var d = -123
	fmt.Printf("%5d, %5d\n", d, d)
	fmt.Printf("%05d, %05d\n", d, d)
	fmt.Printf("%-5d, %-05d\n", d, d)

	var e = -123456789
	fmt.Printf("%5d, %5d\n", e, e)
	fmt.Printf("%05d, %05d\n", e, e)
	fmt.Printf("%-5d, %-05d\n", e, e)
}
