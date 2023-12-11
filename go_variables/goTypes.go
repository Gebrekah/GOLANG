package main

import "fmt"

func main() {
	var intVar1 int8 = -89
	//var intVar2 int16 = 100000 //error overflow
	//var intVar3 uint8 = -23 //error unsigned int
	var intVar4 int = 10000000000000 // recommended unless otherwise the other are needed
	fmt.Println(intVar1, intVar4)

	var intVar5 = 10000000000000 
	intVar6 := 1990293820
	fmt.Println(intVar5, intVar6)
	
	// Floating data types

	var floatVar float32= 1.23456789999 // result: 1.2345679 #7 decimal digit
	var floatVar1 float64 = 1.234567899996777777766655544 // result: 1.2345678999967777 #15 decimal digits
     
	fmt.Println(floatVar, floatVar1)

	var floatVar2 = 1.562562626266
	floatVar3 := 1.562562626266
	fmt.Println(floatVar2, floatVar3)
    
	/* 
	  When you don't declare any type explicitly, the type inferred is `float64` (The default type for float)
	*/

	// Boolean data types
	var boolVar bool = true
	var boolVar1 bool  // if not initialized boolVar1 = false

	fmt.Println(boolVar, boolVar1) 

}