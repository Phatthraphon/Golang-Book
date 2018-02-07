package main

import "fmt"

func main() {
	fmt.Println(">>>>>>>>>>> g <<<<<<<<<<<")
	backticks := `HELLO WORLD!,
Today is good day.`
	fmt.Println(backticks)
	doubleQuotes := "HELLO WORLD!, \nToday is good day."
	fmt.Println(doubleQuotes)	
}