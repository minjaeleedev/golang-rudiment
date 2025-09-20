package main

import "fmt"

func complexIf(day int) {
	if day == 1 {
		fmt.Println("첫째 날입니다.")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	} else if day == 2 {
		fmt.Println("둘째 날입니다.")
		fmt.Println("새로운 팀원 면접이 있습니다.")
	} else if day == 3 {
		fmt.Println("셋째 날입니다.")
		fmt.Println("설계안을 확정하는 날입니다.")
	} else if day == 4 {
		fmt.Println("넷째 날입니다.")
		fmt.Println("예산을 확정하는 날입니다.")
	} else if day == 5 {
		fmt.Println("다섯째 날입니다.")
		fmt.Println("최종 계약하는 날입니다.")
	} else {
		fmt.Println("프로젝트를 진행하세요.")
	}
}

func simpleSwitch(day int) {
	switch day {
	case 1:
		fmt.Println("첫째 날입니다.")
		fmt.Println("오늘은 팀미팅이 있습니다.")
	case 2:
		fmt.Println("둘째 날입니다.")
		fmt.Println("새로운 팀원 면접이 있습니다.")
	case 3:
		fmt.Println("셋째 날입니다.")
		fmt.Println("설계안을 확정하는 날입니다.")
	case 4:
		fmt.Println("넷째 날입니다.")
		fmt.Println("예산을 확정하는 날입니다.")
	case 5:
		fmt.Println("다섯째 날입니다.")
		fmt.Println("최종 계약하는 날입니다.")
	default:
		fmt.Println("프로젝트를 진행하세요.")
	}

}

func main() {
	complexIf(3)
	simpleSwitch(3)
}
