// https://pyrasis.com/book/GoForTheReallyImpatient/Unit17/03
package main

import "fmt"

// https://pyrasis.com/book/GoForTheReallyImpatient/Unit17/03
func main() {
	sample17031()
	sample17032()
}

func sample17031() {
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}
}

func sample17032() {
	for i, j := 0, 0; i < 10; i++, j+=2 { // 컴파일 에러. syntax error: unexpected comma, expecting {
		fmt.Println(i, j)
	}
}
