package main

import (
	"fmt"
)

func main(){

	//map[keyType]valueType

	ages := map[string]int{
		"chandan": 25,
		"sushelamma": 67,
	}

	fmt.Println(ages["chandan"], len(ages))

	// make(map[k]V)


	var scores map[string]int

	scores = make(map[string] int)
	scores["math"] = 20


	fmt.Println(scores["math"])

	users := map[string]string{
		"u1": "chandan",
		"u2": "rahul",
		"u3": "harshith",
	}
	fmt.Println(users)

	delete(users, "u2")
	fmt.Println(users)

}