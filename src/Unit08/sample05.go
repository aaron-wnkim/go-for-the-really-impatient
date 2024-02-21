// https://pyrasis.com/book/GoForTheReallyImpatient/Unit08/05
package main

import "fmt"

func main() {
	sample051()
	sample052()
}

func sample051() {
	var r1 rune = '한'
	var r2 rune = '\ud55c'     // 한
	var r3 rune = '\U0000d55c' // 한

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}

func sample052() {
	var r1 rune = "한"  // 컴파일 에러
	var r2 rune = '한글' // 컴파일 에러
	var r3 rune = "한글" // 컴파일 에러

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}
