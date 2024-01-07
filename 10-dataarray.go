package main

import "fmt"

func main() {
	var names [3] string
	names[0] = "Ahmad"
	names[1] = "Shofi"
	names[2] = "Muharom"

	fmt.Println(names[2])

//membuat array secara langsung

	var values = [5]string{
		"80",
		"90",
		"100",
		"Ahmad",
	}

	fmt.Println(values[3]) // hasilnya Ahmad

//function array

}