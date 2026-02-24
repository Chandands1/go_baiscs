package main

import(
	"fmt"
)

func main(){
   fmt.Println(divide(6,6))
}

func divide(a int, b int)(john int, sangam int){
	john = a/b
	sangam = a +b
	//naked return 
	return
}