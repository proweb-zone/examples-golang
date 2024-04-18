package main

import (
	"fmt"
)

func main(){
	nums := []int{3,2,4}
	target := 6
	result := []int{}

	result = twoSum(nums, target)

for _,item := range result {
	fmt.Println(item)
}

}

func twoSum(nums []int, target int) []int {
	twoSum := []int{}

	for index:= range nums {
		for subIndex:= range nums {
			if index != subIndex && nums[index]+nums[subIndex] == target {
				twoSum = append(twoSum, index)
				twoSum = append(twoSum, subIndex)
				break
			}
		}

		if len(twoSum) >= 2 {
			break
		}
	}

	return twoSum
}
