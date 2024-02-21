// https://pyrasis.com/book/GoForTheReallyImpatient/Unit14/01
package main

import (
	"fmt"
	"runtime"
)

func main() {
	sample141()
	sample142()
}

func sample141() {
	fmt.Println("CPU Count : ", runtime.NumCPU())
}

func sample142() {
	fmt.Println("CPU Count : ", runtime.NumCPU())
}
