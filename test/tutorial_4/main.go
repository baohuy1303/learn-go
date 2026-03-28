package main

import "fmt"

func main() {
	var intArr [3]int32 // fixed length - array type

	var intArr2 [3]int32 = [3]int32{1, 2, 3} // init with values inside
	intArr3 := [3]int32{3, 2, 1} // shorthand
	intArr4 := [...]int32{1, 4, 5, 6} // length can be infered based on num of values

	intArr[1] = 123
	fmt.Println(intArr[0]) // Default to int default which is 0
	fmt.Println(intArr[1:3]) // Slicing like python string, index -> index - 1
	fmt.Println(intArr2)
	fmt.Println(intArr3)
	fmt.Println(intArr4)
	// Go array is contiguous in memory (elements in an array is stored next to each other)

	fmt.Println(&intArr[0]) // Get address with &

	/* ------------------ SLICE - ARRAY LIST ------------------ */

	var intSlice = []int32{1, 3, 5} // similar to array list or dynamic array/stacks in python
	fmt.Printf("The length is %v and capacity is %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // appending
	fmt.Printf("\nThe length is %v and capacity is %v\n", len(intSlice), cap(intSlice))

	// len gets the actual values (not nil) in the array, cap gets capacity
	// capacity is all elements in an array, including empty elements

	intSlice2 := []int32{9, 11}
	intSlice = append(intSlice, intSlice2...) // syntax for adding other vars to array
	fmt.Println(intSlice)

	intSlice3 := make([]int32, 3, 8) // make len 3, cap 8
	fmt.Println(intSlice3)

	/* ------------------ MAP ------------------ */
	var map1 map[string]uint8 = make(map[string]uint8) // Key type string, value type unsignedint
	fmt.Println(map1)

	var map2 = map[string]uint8{"Adam": 10, "Eve": 20, "Jack": 5}
	fmt.Println(map2["Adam"])
	fmt.Println(map2["Huy"]) // nil key just returns default of the value type
	delete(map2, "Adam") // built in delete func for a key
	var age, exist = map2["Adam"] // second var is built in, when declared like this, exist is a bool
	// to signal if the key exist in the map or not
	if exist{
		fmt.Println(age)
	}else{
		fmt.Println("Name/key doesn't exist")
	}

	/* ------------------ LOOP ------------------ */
	for name, age := range map2{ // loop through maps easily + enumerate
		fmt.Printf("The name is %v and age: %v\n", name, age) // no order preserved in go
	}

	for i, value := range intArr{
		fmt.Printf("The index is %v and number: %v\n", i, value)
	}

	// GO HAS NO WHILE LOOP - Declare a var and keep track with the for loop and break if

	//Method 1
	var i int = 0
	for {
		if i >= 10{
			break
		}
		fmt.Println(i+i)
		i++
	}

	//Method 2: more traditional way of for loop
	for i:=0; i < 10; i++{
		fmt.Println(i+i)
	}
}