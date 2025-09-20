package main

import "fmt"

func main() {
	temp := 18

	switch true { // switch 로 줄여쓸 수 있음
	case temp < 10, temp > 30: // 만족하지 않음
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20: // 만족. 실행
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요.")
	case temp >= 15 && temp < 25: // 만족하지만 실행하지 않음. 두 번째에서 이미 실행됨
		fmt.Println("야외 활동하기 좋은 날씨입니다.")
	default:
		fmt.Println("따뜻합니다.")
	}
}
