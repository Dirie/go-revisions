package main

import "fmt"

func main() {


//if-else conditions
	fmt.Println("if-else conditions")
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

	//using else if and short statement
	fmt.Println("\nusing else if and short statement")

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

	//using logical operators
	fmt.Println("\nusing logical operators")
    if 8%2 == 0 || 7%2 == 0 {
        fmt.Println("either 8 or 7 are even")
    }

	//using else if ladder
	fmt.Println("\nusing else if ladder")
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}