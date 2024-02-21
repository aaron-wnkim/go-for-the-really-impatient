// https://pyrasis.com/book/GoForTheReallyImpatient/Unit16
package main

import (
	"fmt"
	"os"
)

// https://go.dev/doc/go1.16#ioutil
func main() {
	sample161()
	sample162()
	sample163()
}

func sample161() {
	var b []byte
	var err error

	b, err = os.ReadFile("./hello.txt")
	if err == nil {
		fmt.Printf("%s", b)
	}
}

func sample162() {
	if b, err := os.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b)
	}
}

func sample163() {
	if b, err := os.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b) // 변수 b를 사용할 수 있음
	} else {
		fmt.Println(err) // 변수 err을 사용할 수 있음
	}

	fmt.Println(b)   // 변수 b를 사용할 수 없음. 컴파일 에러
	fmt.Println(err) // 변수 err을 사용할 수 없음. 컴파일 에러
}
