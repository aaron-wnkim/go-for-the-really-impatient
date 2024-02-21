// https://pyrasis.com/book/GoForTheReallyImpatient/Unit15
package main

import "fmt"

func main() {
	sample151()
	sample152()
	sample153()
	sample154()
	sample155()
}

func sample151() {
	i := 10

	if i >= 5 {
		fmt.Println("5 이상")
	}
}

func sample152() {
	i := 10

	if i >= 5 // 컴파일 에러
	{
		fmt.Println("5 이상")
	}
}

func sample153() {
	i := 10

	if i >= 5 // 컴파일 에러
		fmt.Println("5 이상")
}

func sample154() {
	i := 4

	if i >= 5 {
		fmt.Println("5 이상")
	} else {
		fmt.Println("5 미만")
	}
}

func sample155() {
	i := 6

	if i >= 10 {
		fmt.Println("10 이상")
	} else if i >= 5 && i < 10 {
		fmt.Println("5 이상, 10 미만")
	} else if i >= 0 && i < 5 {
		fmt.Println("0 이상, 5 미만")
	}
}
