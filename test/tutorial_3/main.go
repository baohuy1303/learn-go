package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("test") // type enforced
	var a int = 5
	b := 0
	result, remainder, err := intDivision(a, b)
	if err != nil{
		fmt.Println(err.Error())
	}else if remainder==0{ //has to be on same line as bracket when using else
		fmt.Printf("The result of division is %v", result) // v is value
	}else{
		fmt.Printf("The result of division is %v with remainder %v", result, remainder) // v is value
	}

	// has the same operators as normal && || == != <= >= etc...
	// switch has the same syntax as normal: switch on diff operations, switch on 1 var...
	
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

func intDivision(a, b int) (int, int, error){ // Multiple return + types
	var err error // Built in error type, default is nil (null)
	if b == 0{
		err = errors.New("Cannot divide by 0") // import errors package and raise
		return 0, 0, err
	}

	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}