// https://pyrasis.com/book/GoForTheReallyImpatient/Unit22/05
package main

import "fmt"

func main() {
	sample22051()
	sample22052()
	sample22053()
	sample22054()
	sample22055()
}

func sample22051() {
	a := []int{1, 2, 3, 4, 5}

	b := a[0:5] // a의 인덱스 0부터 5까지 참조

	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(b) // [1 2 3 4 5]
}

func sample22052() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[0:3]) // [1 2 3]
	fmt.Println(a[1:3]) // [2 3]
	fmt.Println(a[2:5]) // [3 4 5]
}

func sample22053() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a[:])        // [1 2 3 4 5]
	fmt.Println(a[0:])       // [1 2 3 4 5]
	fmt.Println(a[:5])       // [1 2 3 4 5]
	fmt.Println(a[0:len(a)]) // [1 2 3 4 5]

	fmt.Println(a[3:])  // [4 5]
	fmt.Println(a[:3])  // [1 2 3]
	fmt.Println(a[1:3]) // [2 3]
}

func sample22054() {
	a := [5]int{1, 2, 3, 4, 5} // 배열 선언

	b := a[:2] // 배열 a의 일부를 부분 슬라이스로 참조
	b[0] = 9   // 부분 슬라이스는 참조이므로 a[0], b[0]의 값이 모두 바뀜

	fmt.Println(a) // [9 2 3 4 5]
	fmt.Println(b) // [9 2]
}

func sample22055() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	b := a[0:6:8] // 인덱스 0부터 6까지 가져와서 부분 슬라이스로 만들고 용량을 8로 설정

	fmt.Println(len(b), cap(b)) // 6 8: 길이가 6이며 용량이 8인 슬라이스
	fmt.Println(b)              // [1 2 3 4 5 6]
}
