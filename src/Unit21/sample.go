// https://pyrasis.com/book/GoForTheReallyImpatient/Unit21
package main

import "fmt"

func main() {
	sample211()
}

func sample211() {
	var a [5]int
	a[2] = 7
	fmt.Println(a)
}

func sample212() {
	var a [5]int = [5]int{32, 29, 78, 16, 81} // int형이며 길이가 5인 배열을 선언하고 초기화
	var b = [5]int{32, 29, 78, 16, 81}        // 배열을 선언할 때 자료형과 길이 생략
	c := [5]int{32, 29, 78, 16, 81}           // 배열을 선언할 때 var 키워드, 자료형과 길이 생략
}

func sample213() {
	a := make([]int, 5, 10)

	fmt.Println(len(a)) // 길이는 5
	fmt.Println(cap(a)) // 용량은 10
}

func sample214() {
	a := make([]int, 5, 10) // 길이가 5이면 a[0], a[1], a[2], a[3], a[4]가 생성

	fmt.Println(a[4]) // 0: make 함수를 사용하면 슬라이스의 요소는 모두 0으로 초기화
	fmt.Println(a[5]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러 발생
	fmt.Println(a[8]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러 발생
}
