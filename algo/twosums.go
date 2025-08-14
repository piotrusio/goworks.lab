package algo

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that
they add up to target. You may assume that each input would have exactly one solution, and you may
not use the same element twice. You can return the answer in any order.

Example 1:

	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

	Input: nums = [3,2,4], target = 6
	Output: [1,2]

Example 3:

	Input: nums = [3,3], target = 6
	Output: [0,1]
*/
func TwoSumA(nums []int, target int) []int {
	arr := make([]int, 2)
	for i, v := range nums {
		lookfor := target - v
		arr[0] = i
		for j, w := range nums[i+1:] {
			if w == lookfor {
				arr[1] = i + 1 + j
				return arr
			}
		}
	}
	return arr
}

// hashmap
func TwoSumB(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		lookfor := target - v
		if index, ok := m[lookfor]; ok {
			return []int{i, index}
		}
	}
	return nil
}
