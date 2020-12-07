package main

func minArray(numbers []int) int {
	left, right := 0, len(numbers) - 1
	for left <= right {
		mid := (left + right) / 2
		if  numbers[mid] > numbers[right] {//说明最小值在右边
			left = mid + 1
		} else {//反之则在左边
			right = right - 1
		}
	}
	return numbers[left]
}
