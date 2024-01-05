package main

import "fmt"

func main() {

// //TIPE1
// 	const firstName = "Ahmad"
// 	const lastName = "Shofi"

// 	fmt.Println(firstName) //Ahmad

//TIPE2
	const (
		firstName = "Ahmad Shofi"
		lastName = "Muharom"
	)

	fmt.Println(firstName, lastName)
}

// constant tidak bisa di ubah
// ERROR
// firstName = "Ahmad"
// lastName = "Shofi"