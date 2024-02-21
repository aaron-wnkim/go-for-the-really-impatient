// https://pyrasis.com/book/GoForTheReallyImpatient/Unit13/01
package main

import "fmt"

func main() {
	sample131()
	sample132()
}

func sample131() {
	a := 8 + 10/2
	fmt.Println(a) // 13
}

func sample132() {
	a := (8 + 10) / 2
	fmt.Println(a) // 9
}
