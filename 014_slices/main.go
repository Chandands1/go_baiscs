package main

import (
	"fmt"
)

func main(){
	// common collection type -> slice
	// slice is dynamic and it can grow
   // []type{...}

   results := []string{"chandan", "go lang", "developer"}

   fmt.Println(results, len(results))
   results[0] = "rahul"

   fmt.Println(results)

   var nums []int

   nums = append(nums, 10)
   nums = append(nums, 20, 30)

   fmt.Println(nums)
}
