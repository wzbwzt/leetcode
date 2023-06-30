package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

/*
思路：
`双指针`


复杂度：
空间复杂度: O(1)
/mnt/data/joelWu/leetcode/cha10时间复杂度： O(n^2), n为数组长度
*/

func twoSum(nums []int, target int) (out []int) {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				out = append(out, i, j)
				return
			}
		}
	}
	return
}

/*
思路：
`哈希表`
空间换时间

复杂度：
空间复杂度: O(n)
时间复杂度： O(n), n为数组长度
*/
func twoSum2(nums []int, target int) (out []int) {
	targer_map := make(map[int]int)
	for i, v := range nums {
		index, ok := targer_map[v]
		if ok {
			out = append(out, i, index)
			return
		}
		targer_map[target-v] = i
	}
	return
}
