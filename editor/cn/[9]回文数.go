//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 示例 1:
//
// 输入: 121
//输出: true
//
//
// 示例 2:
//
// 输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3:
//
// 输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。
//
//
// 进阶:
//
// 你能不将整数转为字符串来解决这个问题吗？
// Related Topics 数学
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindromeArray(x int) bool {
	if x < 0 {
		return false
	}

	sum := 0
	tmp := make([]int, 0)
	for x != 0 {
		tmp = append(tmp, x%10)
		x /= 10
		sum++
	}

	for k, _ := range tmp {
		if tmp[k] != tmp[sum-1-k] {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
