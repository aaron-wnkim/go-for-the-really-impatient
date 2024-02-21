// https://pyrasis.com/book/GoForTheReallyImpatient/Unit23/01
package main

import "fmt"

func main() {
	sample23011()
}

func sample23011() {
	solarSystem := make(map[string]float32) // 키는 string, 값은 float32인 맵 생성 및 공간 할당

	solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	fmt.Println(solarSystem["Earth"]) // 365.25641

	fmt.Println(solarSystem["Pluto"]) // 0: 존재하지 않는 키를 조회(명왕성은 행성 지위 상실...)

	value, ok := solarSystem["Pluto"] // 맵에 키가 있는지 검사할 때는 리턴값을 두 개 활용
	fmt.Println(value, ok)            // 0 false: 맵에 키가 없으면 두 번째 리턴값으로 false가 리턴됨

	if value, ok := solarSystem["Saturn"]; ok {
		fmt.Println(value) // 10756.1995
	}
}
