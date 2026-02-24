package main

import ( 
	"fmt"
)

func main(){
	// this is fixed and can not grow 
	var mark [3] int 
	mark[0] = 50
	mark[1] = 54
	mark[2] = 89

	fmt.Println(mark)

	// array literal

	res := [5]int{2,3,4,5,6}

	fmt.Println(res)

	fmt.Println(len(res))


}