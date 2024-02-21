// https://pyrasis.com/book/GoForTheReallyImpatient/Unit09/01
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sample011()
	sample012()
}

func sample011() {
	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println(len(s1)) // 6: UTF-8 인코딩의 바이트 길이이므로 6
	fmt.Println(len(s2)) // 5: 알파벳 5글자이므로 5
}

func sample012() {
	var s1 string = "한글"
	fmt.Println(utf8.RuneCountInString(s1)) // 2: 문자열의 실제 길이를 구함
}
