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
		ret  = make([]int, 2)
		temp = make(map[int]int, len(nums))
	)

	for index, value := range nums {
		temp[value] = index
	}

	for index1, value := range nums {
		if index2, ok := temp[target-value]; ok && index1 != index2 {
			ret[0], ret[1] = index1, index2
			return ret
		}
	}

	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
