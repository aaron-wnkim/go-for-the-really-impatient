// https://pyrasis.com/book/GoForTheReallyImpatient/Unit06
package main

import "fmt"

func main() {
	sample061()
	sample062()
	sample063()
}

func sample061() {
	i := 10

	if i >= 5 {
		fmt.Println("5 이상")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func sample062() {
	// 세미콜론
	fmt.Println("Hello, world!")
	fmt.Println("Hello,")
	fmt.Println("world!")

	// 주석
	// Hello, world!
	// fmt.Println("Hello, world!")

	var sum int = 1 + 2 // 더하기

	fmt.Println("Hello" /*안녕하세요. */)

	/*
		fmt.Println("Hello,")
		fmt.Println("world!")
	*/
}

func sample063() {
	var num int = 1
	if num == 1 {
		fmt.Println(num)
	}
}
