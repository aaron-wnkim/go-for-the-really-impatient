// https://pyrasis.com/book/GoForTheReallyImpatient/Unit22/01
package main

import "fmt"

func main() {
	sample22011()
	sample22012()
}

func sample22011() {
	a := []int{1, 2, 3}

	a = append(a, 4, 5, 6)

	fmt.Println(a) // [1 2 3 4 5 6]

}

func sample22012() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	a = append(a, b...) // 슬라이스 a에 슬라이스 b를 붙일 때는 b...을 씀

	fmt.Println(a) // [1 2 3 4 5 6]
}
