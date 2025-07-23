package main

import (
	"fmt"
	"time"
)

// New DynamicArray
// Get by index return element, error
// Set value on index return error
// Append new element with resize
// Insert element on index
// Delete by index
// Size(), IsEmpty, String

func main() {
	fmt.Println("Hello world")

	var arr0 [3]int
	arr1 := [...]int{1,2,3}
	arr2 := [3]int{1,2,3}
	arr3 := [5]int{1,2}

	fmt.Println("Len of arr0: ",len(arr0))
	fmt.Println("Len of arr1: ",len(arr1))
	fmt.Println("Len of arr2: ",len(arr2))
	fmt.Println("Len of arr3: ",len(arr3))

	// basic if
	if arr1 == arr2 {
		fmt.Println("Equal")
	}

	// basic if-else
	if arr0 == arr1 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
	
	// if-else-if
	if arr1[0] == 3 {
		fmt.Println("Yes print from 1 ", arr1[0])
	} else if arr0[1] == 3 {
		fmt.Println("Yes print from 2 ", arr1[0])
	} else if arr0[2] == 3 {
		fmt.Println("Yes print from 3 ", arr1[0])
	} else {
		fmt.Println("No print failed")
	}

	// if with initialization
	if arr4 := [...]int{1,2,3,4}; len(arr4) == 4 {
		fmt.Println("Array element: ", arr4[3])
	}

	// classic for loop
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("The value of index %d is %d\n", i, arr3[i])
	}

	// while style loop
	var i int
	for i < 10 {
		fmt.Println("While loop i:", i)
		i++
	}

	// loop through collection
	for i, v := range arr2 {
		fmt.Printf("The value of index %d is %d\n", i, v)
	}

	// range only with intex
	for i := range arr3 {
		fmt.Printf("The index is %d\n", i)

	}
	
	// range only with value
	for _, v := range arr3 {
		fmt.Printf("The value is %d\n", v)
	}

	// infinite loop
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("I am in the infinite loop")
	}
}