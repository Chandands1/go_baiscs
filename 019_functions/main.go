package main

import(
	"fmt"
)

func main(){

	sum1 := add(6,5)

	fmt.Println(sum1)

	fmt.Println(sumAndProduct(5,5))

	onlySum, _ := sumAndProduct(7,7)

	fmt.Println(onlySum)

}

func add(a int, b int)int{
	return a+b

}

func sumAndProduct(a int, b int) (int, int){
	sum := a+b
	product := a * b
	return sum, product
}