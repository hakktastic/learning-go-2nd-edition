package main

import "fmt"

func main() {

	// exercise 1
	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i)
	fmt.Println(f)

	// exercise 2
	const value = 10
	ii := value
	ff := float64(value)

	fmt.Println(ii)
	fmt.Println(ff)

	// exercise 3

	// int16
	// –32768 to 32767
	// int32
	// –2147483648 to 2147483647
	// int64
	// –9223372036854775808 to 9223372036854775807
	// uint8
	// 0 to 255
	// uint16
	// 0 to 65535
	// uint32
	// 0 to 4294967295
	// uint64
	// 0 to 18446744073709551615

	var b int8 = 127
	var small1 int32 = 2147483647
	var bigI uint64 = 9223372036854775807

	fmt.Println(b, ", add 1 -> ", b+1)
	fmt.Println(small1, ", add 1 -> ", small1+1)
	fmt.Println(bigI, ", add 1 -> ", bigI+1)
}
