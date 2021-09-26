package main

import "fmt"

func main() {

	// strings

	var name1 string = "this is a string" //strings can't be single quotes
	var name2 = "this is a string with type inference"
	var name3 string // no value provided to the variable, strings have a default value of ""

	fmt.Println(name1, name2, name3)

	// walrus operator

	name4 := "this is a string with walrus operator" // can't be used outside of a function

	fmt.Println(name4)

	// Integers

	var age1 int = 20
	var age2 = 30
	var age3 int // no value provided to the variable, integers have a default value of 0
	age4 := 40

	fmt.Println(age1, age2, age3, age4)

	// bits & memory
	var bit1 int8 = 127 // max value of an int8 is 127
	var bit2 int8 = -128 
	var bit3 uint8 = 255 // no overflows since these are unsigned integers
	bit4 := 123 // type inferred as an int which defaults to compiler's int size

	fmt.Println(bit1, bit2, bit3 , bit4)

	// floats

	var float1 float32 = 25.98
	var float2 float64 = 1232312312312.234343343//max float value is 1.7976931348623157e+308
	float3 := 1.5 // type inferred as float64 always

	fmt.Println(float1, float2, float3)

}
