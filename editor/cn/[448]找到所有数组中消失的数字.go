//给定一个范围在 1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
//
// 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
//
// 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
//
// 示例:
//
//
//输入:
//[4,3,2,7,8,2,3,1]
//[1,2,3,4,8,2,7,8]
//
//输出:
//[5,6]
//
// Related Topics 数组
// 👍 425 👎 0

package cn

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nIdx := int(math.Abs(float64(nums[i]))) - 1
		if nums[nIdx] > 0 {
			nums[nIdx] *= -1
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
