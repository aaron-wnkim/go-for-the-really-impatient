// https://pyrasis.com/book/GoForTheReallyImpatient/Unit08/03
package main

func main() {
	var num1 complex64 = 1 + 2i           // 실수부 1, 허수부 2
	var num2 complex64 = 4.2342 + 2.3527i // 실수부 소수점 사용 4.2342,
	// 허수부 지수 표기법 2.3527
	var num3 complex64 = 1.e+3i       // 실수부 지수 표기법 1.e, 허수부 3
	var num4 complex64 = 7.27151e-10i // 실수부 없음, 허수부 지수 표기법 7.27151e-10

	var num5 complex128 = 1 + 2i                   // 실수부 1, 허수부 2
	var num6 complex128 = 5.32521e-10 + .12345e+2i // 실수부 지수 표기법 5.32521e-10,
	// 허수부 지수 표기법 .12345E+2

	var r1 float32 = real(num1) // num1의 실수부 1
	var i1 float32 = imag(num1) // num1의 허수부 2

	var r2 float64 = real(num5) // mum5의 실수부 1
	var i2 float64 = imag(num5) // num5의 허수부 2

	var num1 complex64 = complex(1, 2)                    // 실수부 1, 허수부 2
	var num2 complex128 = complex(5.32521e-10, .12345e+2) // 실수부 지수 표기법 5.32521e-10, 허수부 지수 표기법 .12345E+2
}
