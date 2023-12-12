package main

import "fmt"

func main() {

	// If control flow statement
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

	// Switch control flow statement

	var num2 = 3

	switch num2 {
	case 1:
		fmt.Println(num2)
	case 2:
		fmt.Println(num2)
	case 3:
		fmt.Println(num2)
	case 4:
		fmt.Println(num2)
	case 5:
		fmt.Println(num2)
	default:
		fmt.Println("Not found")
	}

	// Switch with a short statement

	switch num3 := 7; num3 {
	case 1:
		fmt.Println(num2)
	case 2:
		fmt.Println(num2)
	case 3:
		fmt.Println(num2)
	case 4:
		fmt.Println(num2)
	case 5:
		fmt.Println(num2)
	default:
		fmt.Println("Not found")
	}

	// Combining multiple Switch cases

	switch dayOfWeek := 2; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Working days")
	default:
		fmt.Println("Weekend")
	}

	// Switch with no expression # is the same if else if flow statement

	var num3 int8 = 8
	switch {
	case num3 > 7:
		fmt.Println(num3)
	case num3 <= 7 || num3 > 5:
		fmt.Println(num3)
	default:
		fmt.Println("")
	}

	// For loop

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// infinite loop with break statement
	for {

		fmt.Println("Infinite")
		break
	}
	// While loop using For loop with continue statement

	j := 10

	for j < 25 {
		fmt.Println(j)
		if j == 20 {
			j += 2
			continue
		}
		j += 5

	}

}
