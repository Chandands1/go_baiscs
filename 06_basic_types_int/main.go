package main

import (
	"fmt"
)

func main(){
	views1 := 1000
	views2 := 2000

	totalsViews := views1 + views2

	likes := 10
	likes++
	likes++

	avgViews := totalsViews/2

	fmt.Println(totalsViews)
	fmt.Println(likes)

	fmt.Println(avgViews)
}