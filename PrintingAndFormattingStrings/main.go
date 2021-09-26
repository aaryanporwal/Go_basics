package main

import "fmt"

func main() {
	// Print
	fmt.Print("Hello, ") // no newline char at the end
	fmt.Print("world!\n")
	fmt.Println()

	// Println
	fmt.Println("Hello, world!") // newline char at the end
	fmt.Println()

	// Printf
	firstname:="aaryan"
	lastname:="porwal"
	_float  := 225.555
	fmt.Printf("My name is %v %v \n",firstname, lastname)
	fmt.Printf("My name is %q %q \n",firstname, lastname) // %q adds quotes around the variables
	fmt.Printf("firstname variable has type: %T \n", firstname)
	fmt.Printf("This is a float: %f \n", _float)
	fmt.Printf("This is a float with set precision of 2: %0.2f \n\n", _float)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My firstname is %v and my lastname is %v \n", firstname, lastname)
	fmt.Println(str)
}
