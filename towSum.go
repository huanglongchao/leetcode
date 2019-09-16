package main

/**
	Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
 */
func twoSum(nums []int, target int) []int {
	result := []int{0,0}
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i] + nums[j] == target{
				result = []int{i,j}
			}
		}
	}
	return result
}
