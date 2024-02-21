// https://pyrasis.com/book/GoForTheReallyImpatient/Unit11
package main

func main() {
	sample111()
	sample112()
	sample113()
	sample114()
}

func sample111() {
	const age int = 10
	const name string = "Maria"
	const score int // 컴파일 에러

	age = 20       // 컴파일 에러
	name = "Grace" // 컴파일 에러
}

func sample112() {
	const age = 10       // int
	const name = "Maria" // string
	const address        // 컴파일 에러
}

func sample113() {
	const x, y int = 30, 50       // x = 30, y = 50
	const age, name = 10, "Maria" // age = 10, name = "Maria"
}

func sample114() {
	const (
		x, y      int = 30, 50      // x = 30, y = 50
		age, name     = 10, "Maria" // age = 10, name = "Maria"
	)
}
