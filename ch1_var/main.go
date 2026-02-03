package main

import "fmt"

func main() {
	// fmt.Println("it is working....")
	var message string = "Hello, World!"
	name := "Sam"

	message = fmt.Sprintf("Hi, %v. Welcome %v!", message, name)
	fmt.Println(message)
	//integer numbers
	var num1  int16 = 10
	var num2  int16 = 20
	msgi := fmt.Sprintf("The sum of %v and %v is %v", num1, num2, num1+num2)
	fmt.Println(msgi)

	var a, b int = 5, 15
	sum := a + b
	fmt.Printf("The sum of a= %v and b= %v is %v\n", a, b, sum)

	//floting point numbers
	var numf1  float32 = 10.0
	var numf2  float32 = 20.5
	msgf := fmt.Sprintf("The sum of %v and %v is %v", numf1, numf2, numf1+numf2)
	fmt.Println(msgf)

	var booleanValue bool = true
	fmt.Printf("The boolean value is %v and its type is %T\n", booleanValue, booleanValue)

	//constant
	const pi float64 = 3.14159
	fmt.Printf("The value of pi is %v\n", pi)

}
