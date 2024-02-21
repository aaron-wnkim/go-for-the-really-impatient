// https://pyrasis.com/book/GoForTheReallyImpatient/Unit18
package main

import "fmt"

func main() {
	sample181()
	sample182()
	sample183()
}

func sample181() {
	var a int = 1

	if a == 1 {
		fmt.Println("Error 1") // 에러 1 상황
		return
	}

	if a == 2 {
		fmt.Println("Error 2") // 에러 2 상황
		return
	}

	if a == 3 {
		fmt.Println("Error 1") // 에러 1 상황, 중복 코드
		return
	}

	fmt.Println(a)
	fmt.Println("Success") // 정상 실행

	return
}

func sample182() {
	var a int = 1

	if a == 1 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}

	if a == 2 {
		goto ERROR2 // 에러 2 상황이면 ERROR2 레이블로 이동
	}

	if a == 3 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}

	fmt.Println(a)
	fmt.Println("Success") // 정상 실행

	return

ERROR1: // 에러 처리 1
	fmt.Println("Error 1")
	return

ERROR2: // 에러 처리 2
	fmt.Println("Error 2")
	return
}

func sample183() {
	var a int = 1

	if a == 1 {
		goto ERROR // 여기까지만 실행하고 ERROR 레이블로 이동
		b := 1
		fmt.Println(b)
	}

ERROR:
	fmt.Println("Error")
}
