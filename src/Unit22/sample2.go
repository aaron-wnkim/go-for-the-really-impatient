// https://pyrasis.com/book/GoForTheReallyImpatient/Unit22/02
package main

import "fmt"

func main() {
	sample22021()
	sample22022()
}

func sample22021() {
	a := [3]int{1, 2, 3}
	var b [3]int

	b = a    // 배열의 요소가 모두 복사됨
	b[0] = 9 // b[0]에 9를 대입하면 b의 첫 번째 요소만 바뀜

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [9 2 3]
}

func sample22022() {
	a := []int{1, 2, 3}
	var b []int // 슬라이스로 선언

	b = a    // a를 b에 대입해도 요소가 모두 복사되지 않고 참조만 함
	b[0] = 9 // 슬라이스는 참조이므로 a[0], b[0]의 값이 모두 바뀜

	fmt.Println(a) // [9 2 3]
	fmt.Println(b) // [9 2 3]
}
