// https://pyrasis.com/book/GoForTheReallyImpatient/Unit13
package main

import "fmt"

// https://pyrasis.com/book/GoForTheReallyImpatient/Unit13
func main() {
	sample131()
	sample132()
	sample133()
	sample134()
	sample135()
}

func sample131() {
	var a int = 1
	var b int = 2
	var c int = b
	const d string = "Hello, world!"
}

func sample132() {
	a := 1               // int
	b := 3.5             // float64
	c := "Hello, world!" // string
}

func sample133() {
	a := 1 + 2                // 3: 두 정수 더하기
	b := 2 + 3                // 5: 두 정수 더하기
	c := a + b                // 8: 두 변수 더하기
	d := "Hello, " + "world!" // Hello, world!: 두 문자열 붙이기
}

func sample134() {
	a := 3 - 2 // 1: 두 정수 빼기
	b := 4 - 5 // -1: 두 정수 빼기
	c := a - b // 2: 두 변수 빼기
}

func sample135() {
	a := 2 * 3  // 6: 두 정수 곱하기
	b := 9 * 21 // 189: 두 정수 곱하기
	c := a * b  // 1134: 두 변수 곱하기
}

func sample136() {
	a := 5 / 2  // 2: 두 정수 나누기
	b := 12 / 4 // 3: 두 정수 나누기
	c := a / b  // 0: 두 변수 나누기
}

func sample137() {
	a := 5 % 2 // 1: 5를 2로 나누었을 때 나머지 구하기
}

func sample138() {
	a := 5         // 5
	a += 2         // 7: a에 2를 더한 후 대입
	fmt.Println(a) // 7
}

func sample139() {
	a := "Hello, " // Hello,
	a += "world!"  // Hello, world!: a에 world! 문자열을 붙인 후 대입
	fmt.Println(a) // Hello, world!
}

func sample1310() {
	a := 5         // 5
	a -= 2         // 3: a에 2를 뺀 후 대입
	fmt.Println(a) // 3
}

func sample1311() {
	a := 5         // 5
	a *= 2         // 10: a에 2를 곱한 후 대입
	fmt.Println(a) // 10
}

func sample1312() {
	a := 5         // 5
	a /= 2         // 2: a에 2를 나눈 후 대입
	fmt.Println(a) // 2
}

func sample1313() {
	a := 5         // 5
	a %= 2         // 1: a에서 2를 나눈 후 나머지를 대입
	fmt.Println(a) // 1
}

func sample1314() {
	a := 3                // 00000011
	b := 2                // 00000010
	c := a & b            // 00000010: a와 b의 AND 비트 연산
	fmt.Printf("%08b", c) // 00000010
}

func sample1315() {
	a := 3                // 00000011
	b := 2                // 00000010
	c := a | b            // 00000011: a와 b의 OR 비트 연산
	fmt.Printf("%08b", c) // 00000011
}

func sample1316() {
	a := 3                // 00000011
	b := 2                // 00000010
	c := a ^ b            // 00000001: a와 b의 XOR 비트 연산
	fmt.Printf("%08b", c) // 00000001
}

func sample1317() {
	a := 255              // 11111111
	b := 68               // 01000100
	c := a &^ b           // 10111011: a와 b의 AND NOT 비트 연산
	fmt.Printf("%08b", c) // 10111011
}

func sample1318() {
	a := 3                // 00000011
	b := 2                // 00000010
	a &= b                // 00000010: a와 b를 AND 비트 연산 후 a에 대입
	fmt.Printf("%08b", a) // 00000010
}

func sample1319() {
	a := 3                // 00000011
	b := 2                // 00000010
	a |= b                // 00000011: a와 b를 OR 비트 연산 후 a에 대입
	fmt.Printf("%08b", a) // 00000011
}

func sample1320() {
	a := 3                // 00000011
	b := 2                // 00000010
	a ^= b                // 00000001: a와 b를 XOR 비트 연산 후 a에 대입
	fmt.Printf("%08b", a) // 00000001
}

func sample1321() {
	a := 255              // 11111111
	b := 68               // 01000100
	a &^= b               // 10111011: a와 b를 AND NOT 비트 연산 후 a에 대입
	fmt.Printf("%08b", a) // 10111011
}

func sample1322() {
	a := 7                // 00000111
	b := a << 2           // 00011100: a의 비트를 오른쪽으로 2번 이동
	fmt.Printf("%08b", b) // 00011100
}

func sample1323() {
	a := 112              // 01110000
	b := a >> 3           // 00001110: a의 비트를 왼쪽으로 3번 이동
	fmt.Printf("%08b", b) // 00001110
	```</td>
</tr>
<tr>
  <td>\<\<=</td>
  <td>비트를 왼쪽으로 이동 후 대입</td>
  <td>
현재 변수를 특정 횟수만큼 왼쪽으로 이동한 다음 다시 변수에 대입합니다.

```go
		a := 7                // 00000111
	a <<= 2               // 00011100: a의 비트를 왼쪽으로 2번 이동한 후 a에 대입
	fmt.Printf("%08b", a) // 00011100
}

func sample1324() {
	a := 112              // 01110000
	a >>= 3               // 00001110: a의 비트를 오른쪽으로 3번 이동한 a에 후 대입
	fmt.Printf("%08b", a) // 00001110
}

func sample1325() {
	var a uint8 = 3       // 00000011
	b := ^a               // 11111100: a의 비트를 반전시김
	fmt.Printf("%08b", b) // 11111100
}

func sample1326() {
	a := 3
	b := -2
	c := +a        // a에 양수 부호를 붙임
	d := +b        // b에 양수 부호를 붙임
	fmt.Println(c) // 3: +(3)
	fmt.Println(d) // -2: +(-2)
}

func sample1327() {
	a := 3
	b := -2
	c := -a        // a에 음수 부호를 붙임
	d := -b        // b에 음수 부호를 붙임
	fmt.Println(c) // -3: -(3)
	fmt.Println(d) // 2: -(-2)
}

func sample1328() {
	fmt.Println(1 == 1)             // true: 두 정수가 같으므로 true
	fmt.Println(3.5 == 3.5)         // true: 두 실수가 같으므로 true
	fmt.Println("Hello" == "Hello") // true: 두 문자열이 같으므로 true

	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Println(a == b) // true: 두 배열이 같으므로 true

	c := []int{1, 2, 3}
	fmt.Println(c == nil) // false: 슬라이스를 nil과 비교하여
	// 메모리가 할당되었는지 확인

	d := map[string]int{"Hello": 1}
	fmt.Println(d == nil) // false: 맵을 nil과 비교하여
	// 메모리가 할당되었는지 확인

	e := 1
	var p1 *int = &e
	var p2 *int = &e
	fmt.Println(p1 == p2) // true: 포인터에 저장된 메모리 주소가 같으므로 true
}

func sample1329() {
	fmt.Println(1 != 2)             // true: 두 정수가 다르므로 true
	fmt.Println(3.5 != 5.5)         // true: 두 실수가 다르므로 true
	fmt.Println("Hello" != "world") // true: 두 문자열이 다르므로 true

	a := [3]int{1, 2, 3}
	b := [3]int{3, 2, 1}
	fmt.Println(a != b) // true: 두 배열이 다르므로 true

	c := []int{1, 2, 3}
	fmt.Println(c != nil) // true: 슬라이스를 nil과 비교하여
	// 메모리가 할당되었는지 확인

	d := map[string]int{"Hello": 1}
	fmt.Println(d != nil) // true: 맵을 nil과 비교하여
	// 메모리가 할당되었는지 확인

	e := 1
	f := 1
	var p1 *int = &e
	var p2 *int = &f
	fmt.Println(p1 != p2) // true: 포인터에 저장된 메모리 주소가 다르므로 true
}

func sample1330() {
	fmt.Println(1 < 2)             // true: 1이 2보다 작으므로 true
	fmt.Println(3.5 < 5.5)         // true: 3.5가 5.5보다 작으므로 true
	fmt.Println("Hello" < "world") // true: H가 w보다 ASCII 코드 값이
	// 작으므로 true
}

func sample1331() {
	fmt.Println(2 <= 2)             // true: 2가 2보다 작거나 같으므로 true
	fmt.Println(3.5 <= 5.5)         // true: 3.5가 5.5보다 작거나 같으므로 true
	fmt.Println("Hello" <= "world") // true: H가 w보다 ASCII 코드 값이
	// 작거나 같으므로 true
}

func sample1332() {
	fmt.Println(2 > 1)             // true: 2가 1보다 크므로 true
	fmt.Println(5.5 > 3.5)         // true: 5.5가 3.5보다 크므로 true
	fmt.Println("world" > "Hello") // true: w가 H보다 ASCII 코드 값이 크므로 true

}

func sample1333() {
	fmt.Println(2 >= 2)             // true: 2가 2보다 크거나 같으므로 true
	fmt.Println(5.5 >= 3.5)         // true: 5.5가 3.5보다 크거나 같으므로 true
	fmt.Println("world" >= "Hello") // true: w가 H보다 ASCII 코드 값이
	// 크거나 같으므로 true

}

func sample1334() {
	fmt.Println(2 >= 2)             // true: 2가 2보다 크거나 같으므로 true
	fmt.Println(5.5 >= 3.5)         // true: 5.5가 3.5보다 크거나 같으므로 true
	fmt.Println("world" >= "Hello") // true: w가 H보다 ASCII 코드 값이
	// 크거나 같으므로 true

}

func sample1335() {
	fmt.Println(true && true)   // true: 두 값이 모두 true이므로 true
	fmt.Println(true && false)  // false: 두 값 중 하나가 false이므로 false
	fmt.Println(false && false) // false: 두 값이 모두 false이므로 false

}

func sample1336() {
	fmt.Println(true || true)   // true: 두 값이 모두 true이므로 true
	fmt.Println(true || false)  // true: 두 값 중 하나가 true이므로 true
	fmt.Println(false || false) // false: 두 값이 모두 false이므로 false

}

func sample1337() {
	fmt.Println(!true)  // false: true의 반대는 false
	fmt.Println(!false) // true: false의 반대는 true

}

func sample1338() {
	a := 1
	b := &a        // a의 메모리 주소를 b에 대입
	fmt.Println(b) // 0xc0820062d0 (메모리 주소)

}

func sample1339() {
	a := new(int)
	*a = 1          // a에 저장된 메모리에 접근하여 1을 저장
	fmt.Println(*a) // 1: a에 저장된 메모리에 접근하여 값을 가져옴
}

func sample1340() {
	c := make(chan int)

	go func() {
		c <- 1 // 채널 c에 1을 보냄
	}()

	a := <-c       // 채널 c에서 값을 가져와서 a에 대입
	fmt.Println(a) // 1

}

func sample1341() {
	a := 1
	a++ // 2: 정수 1을 1 증가시켜서 2

	b := 1.5
	b++ // 2.5: 실수 1.5를 1 증가시켜서 2.5

	c := 1 + 2i
	c++ // (2+2i): 복소수 1+2i를 1 증가시켜서 2+2i

	fmt.Println(a) // 2
	fmt.Println(b) // 2.5
	fmt.Println(c) // (2+2i)

}

func sample1342() {
	a := 1
	b := a++ // 컴파일 에러
	c := ++a // 컴파일 에러
	++a      // 컴파일 에러

}

func sample1343() {
	a := 1
	a-- // 0: 정수 1을 1 감소시켜서 0

	b := 1.5
	b-- // 0.5: 실수 1.5를 1 감소시켜서 0.5

	c := 1 + 2i
	c-- // (0+2i): 복소수 1+2i를 1 감소시켜서 0+2i

	fmt.Println(a) // 0
	fmt.Println(b) // 0.5
	fmt.Println(c) // (0+2i)

}

func sample1344() {
	a := 1
	b := a-- // 컴파일 에러
	c := --a // 컴파일 에러
	--a      // 컴파일 에러

}