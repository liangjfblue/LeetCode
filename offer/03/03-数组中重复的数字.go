/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
/**
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


限制：

2 <= n <= 100000
*/
package _3

//借助hashtable来判断是已遍历的元素
//时间复杂度O(n), 空间复杂度O(n)
func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	existM := make(map[int]int)
	for k, num := range nums {
		if _, ok := existM[num]; ok {
			return num
		}
		existM[num] = k
	}

	return -1
}

//利用占屎坑的原理, 对应值放到和值相等的下标的位置, 若已有元素证明是重复元素
//时间复杂度O(n), 空间复杂度O(1)
func findRepeatNumber2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			tmp := nums[i]
			nums[i], nums[tmp] = nums[tmp], nums[i]
		}
	}

	return -1
}
