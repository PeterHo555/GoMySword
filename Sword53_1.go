package main

func search(nums []int, target int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			ans ++
		}
	}
	return ans

}
