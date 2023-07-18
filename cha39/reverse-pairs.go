package main

/*
给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。

你需要返回给定数组中的重要翻转对的数量。

示例 1:

输入: [1,3,2,3,1]
输出: 2
示例 2:

输入: [2,4,3,5,1]
输出: 3
注意:

给定数组的长度不会超过50000。
输入数组中的所有数字都在32位整数的表示范围内。

*/

/*
思路：
`归并排序`
在得到左右有序序列之后，合并左右有序序列之前,统计符合条件的对数

复杂度：
时间复杂度；O(nlogn)
空间复杂度；O(n)
*/

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	mergeSort(nums, &count, 0, len(nums)-1)
	return count

}

func mergeSort(nums []int, count *int, start, end int) {
	if start == end {
		return
	}
	mid := start + (end-start)>>1

	mergeSort(nums, count, start, mid)
	mergeSort(nums, count, mid+1, end)

	// 此时左右序列已升序，现在做：合并前的统计、以及合并
	i := start
	j := mid + 1
	for i <= mid && j <= end { // i j 都不越界
		if nums[i] > 2*nums[j] {
			*count += mid - i + 1 // i 到 mid，都ok
			j++                   // 考察下一个j，继续找 i
		} else { // 当前i不满足，考察下一个i
			i++
		}
	}
	i = start
	j = mid + 1 // 复原 i j 指针，因为现在要合并左右序列

	temp := make([]int, end-start+1) // 辅助数组，存放合并排序的数
	index := 0                       // 从0开始
	for i <= mid && j <= end {       // 如果 i j 都没越界
		if nums[i] < nums[j] { // nums[i]更小
			temp[index] = nums[i] // 取nums[i]，确定了temp[index]
			index++               // 确定下一个
			i++                   // 考察下一个i，j不动
		} else {
			temp[index] = nums[j]
			index++
			j++
		}
	}
	for i <= mid { // 如果 i 没越界，j越界了
		temp[index] = nums[i] // i 和 i右边的都取过来
		index++               // 确定下一个数
		i++
	}
	for j <= end { // j 没越界，i越界了
		temp[index] = nums[j] // j 和 j右边的都取过来
		index++               // 确定下一个数
		j++
	}
	k := 0
	for i := start; i <= end; i++ { // 根据合并后的情况，更新nums
		nums[i] = temp[k]
		k++
	}
}
