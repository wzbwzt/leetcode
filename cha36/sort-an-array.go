package main

/*
给你一个整数数组 nums，请你将该数组升序排列。



示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]


提示：

1 <= nums.length <= 50000
-50000 <= nums[i] <= 50000

*/

/*
思路：
`归并排序`
分治的思想来对序列进行排序。对一个长为 n 的待排序的序列，我们将其分解成两个长度为 n/2
的子序列。每次先递归调用函数使两个子序列有序，然后我们再线性合并两个有序的子序列使整个序列有序.


复杂度：
时间复杂度；O(nlogn)
空间复杂度；O(n)
*/

func sortArray(nums []int) []int {
	merge_sort(nums, 0, len(nums)-1)
	return nums
}

func merge_sort(nums []int, i, j int) {
	if i == j {
		return
	}
	mid := (i + j) / 2
	merge_sort(nums, i, mid)
	merge_sort(nums, mid+1, j)

	l, r := i, mid+1
	tmp := make([]int, 0, j-i+1)
	for l <= mid || r <= j {
		if l > mid || (r <= j && nums[r] < nums[l]) {
			tmp = append(tmp, nums[r])
			r++
		} else {
			tmp = append(tmp, nums[l])
			l++
		}
	}

	for ii := i; ii < j+1; ii++ {
		nums[ii] = tmp[ii-i]
	}
}
