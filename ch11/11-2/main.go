package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func PrintHouseInfo(house House, name string) {
	fmt.Println(name, " info")
	fmt.Printf("Address: %s\n", house.Address)
	fmt.Printf("Size: %d\n", house.Size)
	fmt.Printf("Price: %f\n", house.Price)
	fmt.Printf("Type: %s\n", house.Type)
	fmt.Println()
}
func main() {
	// 초깃값을 생략하면 모든 필드는 기본값
	var house House
	PrintHouseInfo(house, "house")

	var house2 House = House{"서울시 강동구", 28, 9.80, "아파트"}
	PrintHouseInfo(house2, "house2")

	var house3 House = House{
		"서울시 관악구",
		23,
		9.99,
		"빌라", // 마지막 쉼표 필요
	}
	PrintHouseInfo(house3, "house3")

	// 일부 필드만 초기화
	var house4 House = House{Size: 28, Type: "아파트"}
	PrintHouseInfo(house4, "house4")

	// 여러 줄 초기화
	var house5 House = House{
		Size: 28,
		Type: "아파트",
	}
	PrintHouseInfo(house5, "house5")
}
