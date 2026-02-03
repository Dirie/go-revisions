package main

import "fmt"

func main() {

	//looping through for
	fmt.Println("looping through for")
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

	//using for with init, condition, post
	fmt.Println("using for with init, condition, post")

    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

	//using for with range
	fmt.Println("using for with range")
    for i := range 3 {
        fmt.Println("range", i)
    }

	//infinite loop with break and continue
	fmt.Println("infinite loop with break and continue")
    for {
        fmt.Println("loop - 1")
        break
    }

	//using continue
	fmt.Println("\nusing continue")
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}