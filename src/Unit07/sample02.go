package main

import "fmt"

func main() {
	sample1()
	sample2()
	sample3()
}

func sample1() {
	var x, y int = 30, 50       // x = 30, y = 50
	var age, name = 10, "Maria" // age = 10, name = "Maria"

	a, b, c, d := 1, 3.4, "Hello, world!", false // a = 1, b = 3.4, c = "Hello, world!", d = false

	var x, y int
	var age int

	x, y, age = 10, 20, 5 // x = 10, y = 20, age 5: 병렬할당

	var (
		x, y      int = 30, 50      // x와 y는 int형으로 결정
		age, name     = 10, "Maria" // age는 int, name은 string으로 결정
	)
}

func sample2() {
	a := 2
	b := 2

	_ = b // 사용하지 않는 변수로 인한 컴파일 에러 방지

	fmt.Println(a)
}

func sample3() {
	fmt.Println("Hello, world!")
}
