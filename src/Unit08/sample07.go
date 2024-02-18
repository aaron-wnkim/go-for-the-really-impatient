package main

import (
	"fmt"
	"math"
)

func main() {
	sample071()
	sample072()
	sample073()
}

func sample071() {
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16
	var num3 uint32 = math.MaxUint32
	var num4 uint64 = math.MaxUint64

	fmt.Println(num1) // 255
	fmt.Println(num2) // 65535
	fmt.Println(num3) // 4294967295
	fmt.Println(num4) // 18446744073709551615
}

func sample072() {
	var num1 uint8 = math.MaxUint8 + 1   // 오버플로우 컴파일 에러 발생
	var num2 uint16 = math.MaxUint16 + 1 // 오버플로우 컴파일 에러 발생
	var num3 uint32 = math.MaxUint32 + 1 // 오버플로우 컴파일 에러 발생
	var num4 uint64 = math.MaxUint64 + 1 // 오버플로우 컴파일 에러 발생
}

func sample073() {
	var num1 uint8 = 0 - 1  // 오버플로우 컴파일 에러 발생
	var num2 uint16 = 0 - 1 // 오버플로우 컴파일 에러 발생
	var num3 uint32 = 0 - 1 // 오버플로우 컴파일 에러 발생
	var num4 uint64 = 0 - 1 // 오버플로우 컴파일 에러 발생
}
