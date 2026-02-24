package main

import "fmt"

func moveZeroes(nums []int)  {
	if len(nums) == 0 || len(nums) == 1 { 
		return 
	}

	for slow, fast := 0, 0; fast < len(nums); {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]

			slow++
		}
			
		fast++
	}
}

func main() {
	nums := []int{1,2,3,0,1,2}
	moveZeroes(nums)
	fmt.Println(nums)
}