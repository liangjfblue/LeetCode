//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 4146 👎 0

package cn

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	sum := len(s)
	if sum <= 1 {
		return sum
	}

	max := 0
	l, r := 0, 0
	m := make(map[byte]bool)
	for l < sum && r < sum {
		if _, ok := m[s[r]]; !ok {
			m[s[r]] = true
			max = int(math.Max(float64(max), float64(r-l+1)))
			r++
		} else {
			delete(m, s[l])
			l++
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
