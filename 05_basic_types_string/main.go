package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "Chandan" 
	secondName := "DS"

	fullName := firstName + " " + secondName

	fmt.Println(fullName)

	fmt.Println(strings.ToUpper(fullName))

	
}