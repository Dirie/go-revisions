package main

import "fmt"

func main() {


	//arrays
	fmt.Println("arrays")
    var a1 [5]int
    fmt.Println("emp:", a1)

	//set values and get value
	fmt.Println("\nset and get values")
    a1[4] = 100
    fmt.Println("set:", a1)
    fmt.Println("get:", a1[4])

	//len function
	fmt.Println("\nlen function")
    fmt.Println("len:", len(a1))

	//declare and initialize array
	fmt.Println("\ndeclare and initialize array")
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	//array literal
	fmt.Println("\narray literal")
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	//indexed array literal
	fmt.Println("\nindexed array literal")
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

	//multidimensional array
	fmt.Println("\nmultidimensional array")
    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

	//multidimensional array literal
	fmt.Println("\nmultidimensional array literal")
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)


	//multidimensional array literal
	fmt.Println("\nmultidimensional array literal")
    multD := [2][2][3]int{
		{
			{1, 2, 3},
			{4, 5, 6},
		},	
		{
			{7, 8, 9},
			{10, 11, 12},
		},
    }

	fmt.Println(multD[1][0][2]) // 9
    fmt.Println("3d: ", multD)


    var a [2][3][4]int

    // fill with some values
    val := 1
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a[i]); j++ {
            for k := 0; k < len(a[i][j]); k++ {
                a[i][j][k] = val
                val++
            }
        }
    }

    // iterate and print
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(a[i]); j++ {
            for k := 0; k < len(a[i][j]); k++ {
                fmt.Printf("a[%d][%d][%d] = %d\n", i, j, k, a[i][j][k])
            }
        }
    }

	//beautify output.
	fmt.Println("\nbeautify output.")
	    // iterate and print
    for i := 0; i < len(a); i++ {
		fmt.Print("\n[")
        for j := 0; j < len(a[i]); j++ {
			fmt.Print("\n  [")
            for k := 0; k < len(a[i][j]); k++ {
                fmt.Printf(" %d",  a[i][j][k])
            }
			fmt.Print("]")

        }
		fmt.Print("\n]")
	
    }

	//using range to iterate over multidimensional array
	fmt.Println("\nusing range to iterate over multidimensional array")
	

	 h := [2][2][3]int{
        {
            {1, 2, 3},
            {4, 5, 6},
        },
        {
            {7, 8, 9},
            {10, 11, 12},
        },
    }

    for i, plane := range h {          // first dimension
        for j, row := range plane {    // second dimension
            for k, val := range row {  // third dimension
                fmt.Printf("h[%d][%d][%d] = %d\n", i, j, k, val)
            }
        }
    }



}






