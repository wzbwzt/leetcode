package main

import "math"

/*
给定平面上  n 对不同的点，“回旋镖” 是由点表示的元组  (i, j, k) ，其中  i  和  j  之间的距离和  i  和  k  之间的距离相等（需要考虑元组的顺序）。

找到所有回旋镖的数量。你可以假设  n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。

示例:


输入:
[[0,0],[1,0],[2,0]]

输出:
2

解释:
两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]

*/

/*
思路：
假设有m个点到point[i]的位置相同，那么考虑顺序的情况下，从m中选两个的组合为
$A_{m}^{2}=m*(m-1)$

通过哈希表key存储距离,value存储节点数量


复杂度：
时间复杂度:O(n^2)
空间复杂度:O(n)
*/

func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	var out int

	for _, point_i := range points {
		counter := make(map[int]int, 0)
		for _, point_j := range points {
			distance := math.Pow(float64(point_j[0]-point_i[0]), 2) + math.Pow(float64(point_j[1]-point_i[1]), 2)
			counter[int(distance)] += 1
		}

		for _, count := range counter {
			out += count * (count - 1)
		}
	}
	return out
}
