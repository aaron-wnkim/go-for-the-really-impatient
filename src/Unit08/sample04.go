// https://pyrasis.com/book/GoForTheReallyImpatient/Unit08/04
package main

import "fmt"

func main() {
	var num1 byte = 10
	var num2 byte = 0x32
	var num3 byte = 'a'

	var num3 byte = "a"  // 컴파일 에러
	var num3 byte = 'ab' // 컴파일 에러
	var num3 byte = "ab" // 컴파일 에러

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
}
