package main

import (
	"fmt"
)

func main(){
	// this make method is the slice and takes the length and capacity
	//make([]T, len, cap)
	scores := make([]int, 0,5)

	scores = append(scores, 100)
	scores = append(scores, 200, 3000)

	

	fmt.Println(scores, len(scores), cap(scores))

	// now exeeding the capacity

	scores = append(scores, 500,3444,444)

	fmt.Println(scores, len(scores), cap(scores))


	todos := []string{"wake up", "do workout"}

	more := []string{"learn golang"}

	//appenging the one slice to another slice

	todos = append(todos, more...)

	fmt.Println(todos)
}
