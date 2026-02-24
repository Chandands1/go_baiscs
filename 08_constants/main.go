package main

import (
	"fmt"
)

func main(){
	const appName = "Go Basics"

	//typed constants

	const maxUpload int = 25

	const discountedPrice float64 = 44.32


	fmt.Println(appName, maxUpload, discountedPrice)
}