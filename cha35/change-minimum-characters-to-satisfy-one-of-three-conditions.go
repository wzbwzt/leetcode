package main

/*
给你两个字符串 a 和 b ，二者均由小写字母组成。一步操作中，你可以将 a 或 b 中的 任一字符 改变为 任一小写字母 。

操作的最终目标是满足下列三个条件 之一 ：

a 中的 每个字母 在字母表中 严格小于 b 中的 每个字母 。
b 中的 每个字母 在字母表中 严格小于 a 中的 每个字母 。
a 和 b 都 由 同一个 字母组成。
返回达成目标所需的 最少 操作数。



示例 1：

输入：a = "aba", b = "caa"
输出：2
解释：满足每个条件的最佳方案分别是：
1) 将 b 变为 "ccc"，2 次操作，满足 a 中的每个字母都小于 b 中的每个字母；
2) 将 a 变为 "bbb" 并将 b 变为 "aaa"，3 次操作，满足 b 中的每个字母都小于 a 中的每个字母；
3) 将 a 变为 "aaa" 并将 b 变为 "aaa"，2 次操作，满足 a 和 b 由同一个字母组成。
最佳的方案只需要 2 次操作（满足条件 1 或者条件 3）。
示例 2：

输入：a = "dabadd", b = "cda"
输出：3
解释：满足条件 1 的最佳方案是将 b 变为 "eee" 。


提示：

1 <= a.length, b.length <= 105
a 和 b 只由小写字母组成

*/

/*
思路：
第一个条件：枚举，A的最大元素要小于B的最小元素;第二个条件:A,B交换位置
注意：枚举时从b开始，因为a是最小的元素


复杂度：
时间复杂度：O(m+n)
空间复杂度：O(26)

*/

func minCharacters(a string, b string) int {
	countA := [26]int{}
	countB := [26]int{}
	for _, v := range a {
		countA[v-'a']++
	}

	for _, v := range b {
		countB[v-'a']++
	}
	length := len(a) + len(b)
	ans := length
	for i := 0; i < 26; i++ {
		ans = min(ans, length-countA[i]-countB[i])
	}

	for i := 1; i < 26; i++ {
		t := 0
		for j := i; j < 26; j++ {
			t += countA[j]
		}

		for j := 0; j < i; j++ {
			t += countB[j]
		}

		ans = min(ans, t)
	}

	for i := 1; i < 26; i++ {
		t := 0
		for j := i; j < 26; j++ {
			t += countB[j]
		}

		for j := 0; j < i; j++ {
			t += countA[j]
		}

		ans = min(ans, t)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
