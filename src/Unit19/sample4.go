// https://pyrasis.com/book/GoForTheReallyImpatient/Unit19/04
// https://pkg.go.dev/math/rand@go1.20#Seed
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sample19041()
	sample19042()
}

func sample19041() {
	i := 7

	switch { // case에 조건식을 지정했으므로 판단할 변수는 생략
	case i >= 5 && i < 10: // i가 5보다 크거나 같으면서 10보다 작을 때
		fmt.Println("5 이상, 10 미만")
	case i >= 0 && i < 5: // i가 0보다 크거나 같으면서 5보다 작을 때
		fmt.Println("0 이상, 5 미만")
	}
}

func sample19042() {
	rand.NewSource(time.Now().UnixNano()) // 현재 시간으로 Seed 값 설정
	switch i := rand.Intn(10); {          // rand.Intn 함수를 실행한 뒤 i에 대입
	case i >= 3 && i < 6: // i가 3보다 크거나 같으면서 6보다 작을 때
		fmt.Println("3 이상, 6 미만") // 코드 실행
	case i == 9: // i가 9일 때
		fmt.Println("9") // 코드 실행
	default: // 모든 case에 해당하지 않을 때
		fmt.Println(i) // 코드 실행
	}
}
