//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
// 示例:
//
// 给定 nums = [3,2,4], target = 6
//
//因为 nums[1] + nums[2] = 2 + 4 = 6
//所以返回 [1, 2]
//
// Related Topics 数组 哈希表
package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
思路：
1、遍历列表，存入map中
2、遍历列表，遍历元素，用目标target减去当前元素的值作为map的key，若value存在并且两者下标不一致，两者的和就是target
*/
func twoSum(nums []int, target int) []int {
	var (
		val = make(map[int]int)
	)

	for k, v := range nums {
		if _, ok := val[v]; !ok {
			val[v] = k
		}

		if idx, ok := val[target-v]; ok && idx != k {
			return []int{idx, k}
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
