package main

import (
	"fmt"
)

func main() {
	sample021()
	sample022()
}

func sample021() {
	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "Go"

	fmt.Println(s1 == s2)     // true: 두 문자열이 같으므로 true
	fmt.Println(s1 + s3 + s2) // 한글Go한글
	fmt.Println("안녕하세요" + s1) // 안녕하세요한글
}

func sample022() {
	var s1 string = "Hello"

	fmt.Printf("%c\n", s1[1]) // e: 문자열에서 두 번째 문자 e
}
