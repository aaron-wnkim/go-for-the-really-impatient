// https://pyrasis.com/book/GoForTheReallyImpatient/Unit21/02
package main

import "fmt"

func main() {
	sample21021()
}

func sample21021() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a // 배열을 대입하면 배열 전체가 복사됨

	fmt.Println(a) // [1, 2, 3, 4, 5]
	fmt.Println(b) // [1, 2, 3, 4, 5]
}
