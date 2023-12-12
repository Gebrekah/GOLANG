package main

import "fmt"

func main() {

	var num1 int = 10

	if num1%3 == 0 {
		fmt.Println("Multiple of 3")
	} else {
		fmt.Println("Is not Multiple of 3")
	}

	//combining multiple conditions using short circuit operators && and || and short statement

	if contVar := " "; contVar != "A" && contVar == " " {
		fmt.Println("Control flow statement practice")
	}

	// If-Else-If Chain
	var grade float32 = 34

	if grade > 75 {
		fmt.Println("Very good")
	} else if grade >= 65 || grade <= 75 {
		fmt.Println("Good")
	} else if grade >= 55 || grade < 65 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
