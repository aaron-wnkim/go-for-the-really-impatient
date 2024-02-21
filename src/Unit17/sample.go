// https://pyrasis.com/book/GoForTheReallyImpatient/Unit17
package main

import (
	"fmt"
)

func main() {
	sample171()
	sample172()
	sample173()
	sample174()
}

func sample171() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func sample172() {
	for i := 0; i < 5; i++ // 컴파일 에러
	{
		fmt.Println(i)
	}
}

func sample173() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i = i + 1 // i++
	}
}

func sample174() {
	for {
		fmt.Println("Hello, world!")
	}
}
