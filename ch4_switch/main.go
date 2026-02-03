package main

import (
    "fmt"
    "time"
)

func main() {

	//switch statement
	fmt.Println("switch statement.")
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

	//multiple expressions in case
	fmt.Println("\nmultiple expressions in case")

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

	//switch without expression
	fmt.Println("\nswitch without expression")
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

	//type switch
	fmt.Println("\ntype switch")

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
		case float64:
            fmt.Println("I'm float64")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }

    whatAmI(true)
    whatAmI(1)
    whatAmI(10.4)
    whatAmI("hey")
}