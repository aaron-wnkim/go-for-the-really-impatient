// https://pyrasis.com/book/GoForTheReallyImpatient/Unit21/01
package main

import "fmt"

func main() {
	sample21011()
	sample21012()
	sample21013()
	sample21014()
}

func sample21011() {
	a := [5]int{32, 29, 78, 16, 81}

	for i := 0; i < len(a); i++ { // len 함수로 배열의 길이를 구한 뒤 배열의 길이 만큼 반복
		fmt.Println(a[i])
	}
}

func sample21012() {
	a := [5]int{32, 29, 78, 16, 81}

	for i, value := range a { // i에는 인덱스, value에는 배열 요소의 값이 들어감
		fmt.Println(i, value)
	}
}

func sample21013() {
	a := [5]int{32, 29, 78, 16, 81}

	for value := range a { // value에는 값 대신 인덱스가 들어감
		fmt.Println(value)
	}
}

func sample21014() {
	a := [5]int{32, 29, 78, 16, 81}

	for _, value := range a { // 인덱스는 생략, value에 배열 요소의 값이 들어감
		fmt.Println(value)
	}
}
