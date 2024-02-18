package main

import "fmt"

func main() {
	sample061()
	sample062()
	sample063()
}

func sample061() {
	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println(num1 + num2)  // 5
	fmt.Println(num1 - num2)  // 1
	fmt.Println(num1 * num2)  // 6
	fmt.Println(num1 / num2)  // 1
	fmt.Println(num1 % num2)  // 1
	fmt.Println(num1 << num2) // 12
	fmt.Println(num1 >> num2) // 0
	fmt.Println(^num1)        // 252
}

func sample062() {
	var num1 int = 3
	var num2 float32 = 2.2

	fmt.Println(num1 + num2)          // 정수 + 실수이므로 컴파일 에러
	fmt.Println(float32(num1) + num2) // 5.2: 정수를 실수로 변환
	fmt.Println(num1 + int(num2))     // 5: 실수를 정수로 변환하여 0.2를 버림
}

func sample063() {
	var num1 uint16 = 10
	var num2 uint32 = 80000

	fmt.Println(num1 + uint16(num2)) // 14474: uint32에서 uint16으로 변환하면서 넘치는 값을 버림
}
