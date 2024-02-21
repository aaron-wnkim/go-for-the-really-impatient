// https://pyrasis.com/book/GoForTheReallyImpatient/Unit23
package main

import "fmt"

func main() {
	sample231()
	sample232()
}

func sample231() {
	var a map[string]int = make(map[string]int) // make 함수로 키는 string 값은 int인 맵에 공간 할당
	var b = make(map[string]int)                // 맵을 선언할 때 map 키워드와 자료형 생략
	c := make(map[string]int)                   // 맵을 선언할 때 var, map 키워드와 자료형 생략
}

func sample232() {
	a := map[string]int{"Hello": 10, "world": 20}

	b := map[string]int{
		"Hello": 10,
		"world": 20, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	fmt.Println(a["Hello"]) // 10
	fmt.Println(b["world"]) // 10
}
