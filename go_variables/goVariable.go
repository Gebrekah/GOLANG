/* Golang Variables, Zero Values, and Type inference */
package main

import "fmt"

func main() {
	// Declaring variable
	var firstName string
	var lastName string
	var age int
	var monthlySalary float32
	var isMarried bool
	/*
		// Declare multiple variables at once
		var (
			firstName string
			age int
			monthlySalary float32
			isMarried bool
		)
	
		// Combine multiple variable declarations of the same type with comma -
		
		var (
			firstName, lastName string // var s, k int, string #Gives error
			age int
			monthlySalary float32
			isMarried bool
		)
	*/
	/* Any variable declared without an initial value will have a 
	zero-value depending on the type of the variable- */

	fmt.Println(firstName, lastName, age, monthlySalary, isMarried) // result: "" "" 0 0 false

	// Declaring Variables with initial Value

	var name string = "GOLANG"
	var isOpenSource bool = true
	var createdBy string = "GOOGLE"
	var createdYear int = 2009

	fmt.Println(name, isOpenSource, createdBy, createdYear)
	
	
}