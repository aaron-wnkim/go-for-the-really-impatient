// https://pyrasis.com/book/GoForTheReallyImpatient/Unit19/03
package main

import "fmt"

// https://pyrasis.com/book/GoForTheReallyImpatient/Unit19/03
func main() {
	sample19031()
}

func sample19031() {
	i := 3

	switch i {
	case 2, 4, 6: // i가 2, 4, 6일 때
		fmt.Println("짝수")
	case 1, 3, 5: // i가 1, 3, 5일 때
		fmt.Println("홀수")
	}
}
