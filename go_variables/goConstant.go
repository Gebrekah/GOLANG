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
}