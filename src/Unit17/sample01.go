// https://pyrasis.com/book/GoForTheReallyImpatient/Unit17/01
package main

import "fmt"

func main() {
	sample17011()
	sample17012()
	sample17013()
}

func sample17011() {
	i := 0
	for {
		if i > 4 {
			break
		}

		fmt.Println(i)
		i = i + 1
	}
}

func sample17012() {
Loop: // Loop 레이블을 지정
	for i := 0; i < 3; i++ { // 반복문 1
		for j := 0; j < 3; j++ { // 반복문 2
			if j == 2 { // j가 2일 때
				break Loop // 중첩된 반복문을 빠져나옴
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}

func sample17013() {
Loop:
	fmt.Println("begin for loop") // 들어가면 안 되는 코드
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop // 컴파일 에러. invalid break label Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}
