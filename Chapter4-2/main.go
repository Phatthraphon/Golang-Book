package main

import "fmt"

func main() {
	/*  fmt.Print("Enter a number : ")
		var input float64
		fmt.Scanf("%f", &input)
		output1 := input * 2
		fmt.Println(output1)
	*/

	fmt.Print("Enter Fahrenheit : ")
	var F float64
	fmt.Scanf("%f", &F)
	Celsius := (F-32)*5/9
	fmt.Printf("Celsius : %.2f", Celsius)
}
