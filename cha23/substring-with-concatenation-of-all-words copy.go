package main

/*
https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：
输入：
s = "barfoothefoobarman",
words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
s = "wordgoodgoodgoodbestword",
words = ["word","good","best","word"]
输出：[]

*/

/*
思路：
`哈希表` `双指针`
s中截取m*k长度的子串
按照words中每个单词的长度切割s，得到子字串




复杂度：
时间复杂度: O(n∗m∗k)  n: 字符串 S 长度, m:words 数组元素个数, k: 为单个 word 字串长度。
空间复杂度: O(m)
*/

func findSubstring(s string, words []string) (out []int) {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	n := len(s)
	m := len(words)
	k := len(words[0])
	if n < m*k {
		return nil
	}
	var word_map = make(map[string]int, len(words))
	for _, word := range words {
		word_map[word]++
	}
	for i := 0; i < n-m*k+1; i++ {
		sub_s := s[i : m*k+i]
		var sub_s_map = make(map[string]int, m)
		for j := 0; j < m*k; j += k {
			w := sub_s[j : j+k]
			_, ok := word_map[w]
			sub_s_map[w]++
			if !ok {
				goto lab
			}
		}
		for k, v := range word_map {
			if sub_s_map[k] != v {
				goto lab
			}
		}

		out = append(out, i)
	lab:
	}
	return
}
