package main

import "fmt"

func main() {
	
	// // TIPE1
	// 	var name = "Ahmad Shofi"
	// 	fmt.Println(name)

	// 	name = "Shofi Muharom"
	// 	fmt.Println(name)

	// // TIPE2
	// 	name := "Ahmad Shofi"
	// 	fmt.Println(name)

	// 	name = "Shofi Muharom" // yang ini tidak boleh diganti menjadi ":="
	// 	fmt.Println(name)

	// TIPE3
		var (
			firstName = "Ahmad"
			middleName = "Shofi"
			lastName = "Muharom"
		)

		fmt.Println(firstName, middleName, lastName)
}


