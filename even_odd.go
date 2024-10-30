package main

import "fmt"

func main(){
	nums := []int {1,2,3,4,5,6,7,8,9,10}

	for i, _ := range(nums) {
		if (nums[i]%2 == 1) {
			fmt.Printf("%d is odd.\n", nums[i])
		} else {
			fmt.Printf("%d is even.\n", nums[i])
		}
			
			
	}
}