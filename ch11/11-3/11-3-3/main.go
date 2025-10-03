package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User  // Level 필드를 갖는 구조체. embedded
	Price int
	Level int // VIPUser의 Level 필드
}

func main() {
	user := User{"송하나", "hana", 23, 10}
	vip := VIPUser{
		// User 구조체 초기화 포함
		User{"화랑", "hwarang", 40, 10},
		250,
		3,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d 레벨: %d\n", user.Name, user.ID, user.Age, user.Level)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 가격: %d만 원 VIP 레벨: %d 유저 레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Price,
		vip.Level,      // VIPUser의 Level
		vip.User.Level, // 포함된 구조체이름을 쓰고 접근
	)
}
