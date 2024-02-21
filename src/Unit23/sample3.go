// https://pyrasis.com/book/GoForTheReallyImpatient/Unit23/03
package main

import "fmt"

func main() {

}

func sample23031() {
	a := map[string]int{"Hello": 10, "world": 20}

	delete(a, "world") // 맵 a에서 world 키 삭제

	fmt.Println(a) // map[Hello:10]
}
