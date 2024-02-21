// https://pyrasis.com/book/GoForTheReallyImpatient/Unit22
package main

func main() {
	sample221()
	sample222()
}

func sample221() {
	var a []int = make([]int, 5) // make 함수로 int형에 길이가 5인 슬라이스에 공간 할당
	var b = make([]int, 5)       // 슬라이스를 선언할 때 자료형과 [] 생략
	c := make([]int, 5)          // 슬라이스를 선언할 때 var 키워드, 자료형과 [] 생략
}

func sample222() {
	a := []int{32, 29, 78, 16, 81} // 슬라이스를 생성하면서 값을 초기화

	b := []int{
		32,
		29,
		78,
		16,
		81, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}
}
