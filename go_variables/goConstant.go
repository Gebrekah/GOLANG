package main

import "fmt"

func main() {

	const constVar int32 = 345
	// constVar = 5643 // cannot assign to constVar (neither addressable nor a map index expression)
	// Multple variable declaration
	/* 
	missing init expr for constVar3
	const (
		constVar1 float32  
		constVar2 string
		constVar3 bool
	)
    */
	const (
		constVar1 float32  = 5.2
		constVar2 string = "Strings"
		constVar3 bool = true
	)
	
	fmt.Println(constVar, constVar1, constVar2, constVar3)
	/*
	var a int = 3
	var c = a + constVar // invalid operation: constVar + a (mismatched types int32 and int) 
	*/
	const a = 5 + 7.5 // Valid 12.5
	const b = 12/5    // Valid 2
	const c = 'z' + 1 // Valid treated as rune (single quatation character)

	//const d = "Hey" + true // Invalid (untyped string constant and untyped boolean constant are not compatible with each other)

	fmt.Println(a, b, c)

	
}