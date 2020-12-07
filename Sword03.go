package main

func findRepeatNumber(nums []int) int {
	myMap := make(map[int]int)
	for _, value := range nums{
		if _, ok := myMap[value]; ok {
			// 存在该数值
			return value
		}else {
			// 第一次出现，标记为1
			myMap[value] = 1
		}
	}
	return -1
}