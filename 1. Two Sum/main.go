package main

func main(){
	nums := []int{3,3,4}
	target := 6

	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	reBuildNums := map[int]int{}
	result := []int{}

	for i, num:= range nums {
		reBuildNums[num] = i
	}

	for i, num:= range nums {
		compareIndex, ok := reBuildNums[target - num]

		if ok && i != compareIndex {
			result = append(result, i)
			result = append(result, compareIndex)
			break
		}
	}

	return result
}
