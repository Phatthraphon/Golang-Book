package main

import "fmt"

func main() {
	//long Declaration
	var x string = "HELLO WORLD"
	fmt.Println(x)
	fmt.Printf("Type: %T\n", x)

	var y string
	y = "HELLO WORLD"
	fmt.Println(y)
	fmt.Printf("Type: %T\n", y)

	//short Declaration
	//type Interface
	z := "HELLO WORLD"
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)

}