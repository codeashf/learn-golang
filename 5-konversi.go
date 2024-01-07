package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64) // 32768
	fmt.Println(nilai16) // -32768 (nilai max 32767)

	var name = "Ahmad Shofi"
	var e = name[7] // e adalah uint8
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}