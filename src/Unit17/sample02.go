// https://pyrasis.com/book/GoForTheReallyImpatient/Unit17/02
package main

import "fmt"

// https://pyrasis.com/book/GoForTheReallyImpatient/Unit17/02
func main() {

}

func sample17021() {
	for i := 0; i < 5; i++ {
		if i == 2 { // i가 2일 때
			continue // 아래 부분 코드를 실행하지 않고 넘어감
		}

		fmt.Println(i)
	}
}

func sample17022() {
Loop:
	for i := 0; i < 3; i++ { // 반복문 1
		for j := 0; j < 3; j++ { // 반복문 2
			if j == 2 { // j가 2일 때
				continue Loop // 아래 부분 코드를 실행하지 않고 반복문 1부터 이어서 실행
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}

func sample17023() {
Loop:
	fmt.Println("begin for loop") // 들어가면 안 되는 코드
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				continue Loop // 컴파일 에러. invalid continue label Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}
