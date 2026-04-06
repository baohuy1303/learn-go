package main

import "fmt"

// POINTERS

func main(){
	var p * int32 = new(int32) // * (star) declares a pointer
	// A pointer will point to a memo location, and that memo location can store another address which is
	// p's value address. For ex: p -> nil -> memoLocation (org).
	// If we init p with int32, which default is 0. It turns into: p -> memo_loc_of_0 -> memoLocation (org)

	fmt.Printf("The value of p (not address) is: %v, and the address is: %v\n", *p, p)
	*p = 10 // sets acutual p value
	fmt.Printf("The value of p (not address) is: %v, and the address is: %v\n\n", *p, p) // same address

	var i int32
	p = &i // sets p to the memo address of i, which means p now stores i's memo address
	// &: gets the memo value of a var

	fmt.Printf("p: %v , i: %v\n", p, i) // see how it's 4 bytes after the first address
	// cuz it's at i's memo address
	*p = 1
	fmt.Printf("p (val): %v , i: %v\n", *p, i)
	//Because p at i's memo, if p change, i change
	// THIS EXPLAINS how slices type change operates, it references pointers not copying values
	
	// THIS IS KEY IMPORTACE FOR FUNCS
	// Function takes in parameters. If we take a var param, then it will create a copy of that var, wasting space
	// If we take in pointer to that var as a param, we only use 1 var
	// Because the param is the pointer to the var that we've created before, saving space.


}