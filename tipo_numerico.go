package main

import (
	"fmt"
	"runtime"
)

// uint8 				8-bit unsigned integer (0 to 255)
// uint16 			16-bit unsigned integer (0 to 65535)
// uint32 			32-bit unsigned integer (0 to 4294967295)
// uint64 			64-bit unsigned integer (0 to 18446744073709551615)
// int8 				8-bit signed integer (-128 to 127)
// int16 				16-bit signed integer (-32768 to 32767)
// int32 				32-bit signed integer (-2147483648 to 2147483647)
// int64 				64-bit signed integer (-9223372036854775808 to 9223372036854775807)
// uint 				deprecated: same as uint64
// int 					deprecated: same as int64
// rune 				deprecated: alias of int32
// byte 				deprecated: alias of uint8
// float32 			32-bit IEEE-754 floating-point number
// float64 			64-bit IEEE-754 floating-point number
// complex64 		complex number with float32 real and imaginary parts
// complex128 	complex number with float64 real and imaginary parts

func main() {
	a := "e"
	b := "éãot"
	c := "u9999"

	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n", d, e, f)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
