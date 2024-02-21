// https://pyrasis.com/book/GoForTheReallyImpatient/Unit22/03
package main

import "fmt"

func main() {
	sample22031()
	sample22032()
}

func sample22031() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3) // make 함수로 공간을 할당

	copy(b, a) // 슬라이스 a의 요소를 슬라이스 b에 복사

	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(b) // [1 2 3]: 슬라이스 b의 길이는 3이므로 a의 요소 3개만 복사됨
}

func sample22032() {
	a := []int{1, 2, 3}
	b := make([]int, 3)

	copy(b, a) // 슬라이스를 복사하였으므로
	b[0] = 9   // b[0]만 바뀌고, a[0]은 바뀌지 않음

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [9 2 3]
}
