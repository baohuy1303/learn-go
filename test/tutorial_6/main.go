package main

import "fmt"

// Define my own type
type carEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner // refernce other structs
	int // if put in type directly, name = type
}

type owner struct{
	name string
	age int
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

// INTERFACE - Parent of different structs/types
type engines interface{
	milesLeft() uint8 // Method signature - will be checked at runtime
}

// This method is attached to carEngine explicitly, declare like this
// Has signature so each struct that wants to use the interface has to define the signature

func (e carEngine) milesLeft() uint8{
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8{
	return e.mpkwh * e.kwh
}

// This is not
func milesLeft_Standalone(e carEngine) uint8{
	return e.gallons * e.mpg
}

// Go will check if the type has the signature at runtime.
// ex: if pass in carEngine, will check if has milesLeft since milesLeft
// is the signature of engine interface
func canGo(e engines, milesLeft uint8) bool{
	if milesLeft > e.milesLeft(){ //.milesLeft() is in the func signature
		return false
	}else{
		return true
	}
}

func main(){
	var myEngine carEngine = carEngine{25, 15, owner{"Huy", 18}, 5} // follow the ordering
	myEngine.gallons = 20
	myEngine = carEngine{mpg: 12, gallons: 123, ownerInfo: owner{"Huy", 18}, int: 5} // define exactly
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo, myEngine.ownerInfo.name, myEngine.ownerInfo.age,
	myEngine.int)

	// Anonymous struct - no name
	var anotherCarEngine = struct{
		mpg uint8
		gallons uint8
	}{10, 5}

	var myElectricEngine electricEngine = electricEngine{15, 10}

	fmt.Println(anotherCarEngine)
	fmt.Println(myEngine.milesLeft()) // Func attached to struct
	fmt.Println(myElectricEngine.milesLeft())
	fmt.Println(milesLeft_Standalone(myEngine)) // Standalone func that any can call
	fmt.Println(canGo(myEngine, 200))
}