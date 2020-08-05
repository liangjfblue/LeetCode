/**
 *
 * @author liangjf
 * @create on 2020/7/30
 * @version 1.0
 */
/**
输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。


示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100
*/
package _2

import "math"

//滑动窗口
func maxSubArray(nums []int) int {
	sum := 0
	max := -math.MaxUint32
	for i := 0; i < len(nums); i++ {
		//当前sum小于0, 更新sum为当前值, 因为假如sum加上负值会更小, 加上正值更好, 更大了
		if sum < 0 {
			sum = nums[i]
		} else {
			//当前sum大于0, 不管当前元素是正数还是负数都累加, 因为下面会和max比较才判断是否更新max为sum
			sum += nums[i]
		}

		if sum > max {
			max = sum
		}
	}
	return max
}

//动态规划
//dp数组最后一位是前面的子数组的和
//不关心哪些连续子数组最大, 只关心连续最大值
//因此, dp最后一位加上nums当前总和大于dp最后一位, 就更新dp最后一位的值, 否则不更新;
//每次对比dp最后一位的值和max比较, 若大于max就更新max
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, 0)
	dp = append(dp, nums[0])
	max := -math.MaxUint32

	for i := 0; i < len(nums); i++ {
		//不关心哪些连续子数组最大, 只关心连续最大值
		dp = append(dp, int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i]))))
		max = int(math.Max(float64(dp[i]), float64(max)))
	}
	return max
}
