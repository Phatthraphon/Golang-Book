package main

import "fmt"

func main() {  
    for i := 1; i <= 100; i++ {
        if i%15 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
	}
}
func fizzbuzz(max int) {
	for i := 1; i < max; i++ {
		switch {
		case (i%3 == 0) && (i%5 == 0):
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func mainfizbuzz() {
	fizzbuzz(100)
}

func switchcase() {
	for M :=1; M <=5 ; M++{
		switch M {
		case 0:
		fmt.Println("Zero")
		case 1:
		fmt.Println("One")
		case 2:
		fmt.Println("Two")
		case 3:
		fmt.Println("tree")
		case 4:
		fmt.Println("Four")
		case 5:
		fmt.Println("Five")
		default:
			fmt.Println("Error")
		}
	}
}