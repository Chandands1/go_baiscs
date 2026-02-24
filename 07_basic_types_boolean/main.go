package main

import(
	"fmt"
)

func main(){
	isLogged := false
	isAdmin := true
	hasSubscription := false
	

	// And &&

	canOpenDashboard := isLogged && hasSubscription

	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println(canOpenDashboard)
	fmt.Println(canDeletePost)

	age := 17

	isAdult := age > 18

	fmt.Println(isAdult)
}
