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

	// Complex datatype
	var complexVar complex64 = 6697.3 + 5231.6786666666i // both real and imaginary parts are of float32 type.

	var complexVar1 complex128 = 567.87 + 213.555i // both real and imaginary parts are of float64 type

	var complexVar2 = 32.5 + 67.56i  // Type inferred as `complex128`

	fmt.Println(complexVar, complexVar1, complexVar2)

	// creating a complex number with variables instead of literals

	var realNum = 45.35
	var imginaryNum = 32.75

	// var complexNum = realNum + imginaryNum*i // error use complex keyword
	var complexVar3 = complex(realNum, imginaryNum)

	fmt.Println(complexVar3)

	// Both real and imaginary parts must be of the same floating-point type
   /*
	var realNum1 complex64 = 45.35
	var imginaryNum1 complex128 = 32.75

	var complexVar3 = complex(realNum1, imginaryNum1) // error
	*/
	
	// string datatypes

	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var strVar string = "Strings in Golang are declared either \t using double quotes or back ticks\n"

	// Raw String (Can span multiple lines and displayed as it is. Escape characters are not interpreted)
	var strVar1 string = `Strings in Golang are declared \b either \n
    using double quotes or back ticks` // the space from the start counted here and Escape characters don't have any special meaning in raw strings.

	fmt.Println(strVar, strVar1)

	// Type conversion
 	/*
	var a int64 = 4
	var b int = a  // Compiler Error (Cannot use a (type in64) as type int in assignment)

	var c int = 500

	var result = a + c // Compiler Error (Invalid Operation: mismatched types int64 and int)
	*/

	var a int64 = 4
	var b int = int(a)  // Explicit Type Conversion

	var c float64 = 6.5

	// Explicit Type Conversion
	var result = float64(b) + c  // Works

	fmt.Println(a, b, c, result)

}