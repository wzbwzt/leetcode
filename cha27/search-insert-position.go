package main

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

```
示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0

*/

/*
思路：
`双指针`


复杂度：
时间复杂度：O(n)
空间复杂度：O(1)
*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] < target {
			slow++
		}
		if nums[fast] == target {
			return fast
		}
		fast++
	}

	return slow
}

/*
思路：
`二分`

复杂度：
时间复杂度：O(logn)
空间复杂度：O(1)
*/
func searchInsert2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (right-left)>>1 + left
		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left + 1
}
