package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 100 //var name type =...
	fmt.Println(intNum)

	var float float32 = 55.23 // only 2 types of floats
	fmt.Println(float)

	// have different types like int32 int64 uint8 etc... 
	// think what types to use for best memory and accuracy

	var result float32 = float32(intNum) + float // enforce same type for arithmetic
	fmt.Println(result)

	// division int will be floored (same as java) modulo % gets remainder

	var myString string = "Hello World" //double quotes for 1 line, back tick for multi lines
	var myString2 string = `Hello
world`

	fmt.Println(myString + " " + myString2)
	fmt.Println(len(myString)) // ASCII takes 1 byte
	fmt.Println(len("γ")) // Other takes 2, len func counts bytes not chars
	fmt.Println(utf8.RuneCountInString("γ")) // This package actually counts chars

	var myRune rune = 'a' // rune is another data type that represents chars
	fmt.Println(myRune) //runes are weird (this prints 97)

	var defaultInt int // can declare and Go sets default values (int 0, string '', bool false)
	var boolean bool
	fmt.Println(defaultInt)
	fmt.Println(boolean)

	var autoSets = "type"
	myText := "text" // AUTO sets type from Go based on value
	fmt.Println(autoSets + " " + myText)

	const myConst string = "Cant change" // cant change val of const, and needs to be init with value
	// cant be left empty

}