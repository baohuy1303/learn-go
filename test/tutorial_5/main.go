package main

import (
	"fmt"
	"strings"
)

func main() {
	// GO uses bytes to represent strings
	// If char not ASCII then gonna take 2 bytes
	// résumé means r: 1 byte, é: 2 bytes, ... [r, é] each element takes number of spaces
	// in the array based on their byte size. Ex: r takes 1 space, é takes 2 (2 bytes)

	var myString string = "résumé"
	var indexed = myString[1] // Return unsigned int 8
	fmt.Printf("value: %v, type: %T\n", indexed, indexed)
	for i, v := range myString { // skips index 2 because é takes 2 bytes => 1, 2
		fmt.Println(i, v)
	}
	fmt.Println(len(myString)) // Prints out number of bytes

	//-------------- EASIER way to deal with strings: cast array of rune --------------

	var myString_Rune = []rune("résumé")
	var indexed_Rune = myString_Rune[1] // Return type int32
	fmt.Printf("value: %v, type: %T\n", indexed_Rune, indexed_Rune)
	for i, v := range myString_Rune{ // No byte skipping
		fmt.Println(i, v)
	}
	fmt.Println(len(myString_Rune)) // Correct len

	var stringAdd = []string{"h", "e", "l", "l", "o"} 
	var finalString string = ""
	for i := range stringAdd{
		finalString += stringAdd[i] // Creates a new string everytime => inefficient
	}
	fmt.Println(finalString)
	// string is immutable

	//More efficient use stringBuilder
	var stringSlice = []string{"h", "e", "l", "l", "o"}
	var stringBuilder strings.Builder
	for i := range stringSlice{
		stringBuilder.WriteString(stringSlice[i]) // Built-in builder, add each char to stringBuilder
	}
	var castStr = stringBuilder.String()
	fmt.Println(castStr)

}