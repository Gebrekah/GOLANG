package main

import "fmt"

func main() {
	// Array declaration
	var arrayVar [5]int
	var arrayVar1 [10]string
	var arrayVar2 [2]bool
	var arrayVar3 [4]complex128
	var arrayVar4 [5]float32
	fmt.Println(arrayVar, arrayVar1, arrayVar2, arrayVar3, arrayVar4)

	// Accessing an array
	arrayVar[2] = 3
	//arrayVar[10] = 10 //invalid argument: index 10 out of bounds [0:5]
	fmt.Println(arrayVar[1], arrayVar1[0], arrayVar2[1], arrayVar3[3], arrayVar4[2])

	// Declaring and initializing an array at the same time
	var num = [5]int{2, 4, 6, 8, 10}
	fmt.Println(num)
	num1 := [5]int{2, 5, 9, 8, 10}
	fmt.Println(num1)

	// You don't need to initialize all the elements of the array.
	// The un-initialized elements will be assigned the zero value of the corresponding array type
	num2 := [5]int{2}
	fmt.Println(num2)

	// Letting Go compiler infer the length of the array
	num3 := [...]int{3, 5, 7, 9, 11, 12, 17}
	fmt.Println(num3)

}
