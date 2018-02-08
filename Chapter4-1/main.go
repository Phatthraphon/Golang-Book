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

	//const
	//const t string = "HELLO WORLD"
	//t = "Other String"

	var (
		a = 5
		b =10
		c = 15
	)

	fmt.Println(a,b,c)
	fmt.Println(a+b+c)

	v1,v2  := "first","sec"
	v1,v2 = v2,v1
	fmt.Println(v1)
	fmt.Println(v2)

}
