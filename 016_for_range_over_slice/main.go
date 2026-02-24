package main

import( 
	"fmt"
)

func main(){

	view := []int{10,20,30,40,50,60}

	//for range
	total := 0

	for i, v := range view{
		fmt.Println("day", i, "views: ", v)
		total = total +v
	}

	fmt.Println(total)
	

}
