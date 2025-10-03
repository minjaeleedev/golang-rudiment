package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House
	house.Address = "서울"
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억 원\n", house.Price)
	fmt.Println("타입:", house.Type)
}
