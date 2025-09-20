package main

import "fmt"

type ColorType int // 별칭 ColorType 선언하고 const 열거값 정의
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

// ColorType 열거값에 따라 문자열을 반환
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
