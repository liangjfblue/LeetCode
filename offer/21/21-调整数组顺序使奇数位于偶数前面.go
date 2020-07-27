/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。


示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。


提示：

1 <= nums.length <= 50000
1 <= nums[i] <= 10000
*/
package _1

//双指针法:
//找左右不符合的首对元素, 交换元素, 并左右向中间逼近
//若右边是偶数, 移动右指针
//若左边是奇数, 移动左指针
//时间复杂度O(N)，空间复杂度O(1)
func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	l := 0
	r := len(nums) - 1

	for l < r {
		if nums[l]%2 == 0 && nums[r]%2 == 1 {
			nums[l], nums[r] = nums[r], nums[l]
			r--
			l++
		} else if nums[r]%2 == 0 {
			//左边有偶数, 移动右边找奇数
			r--
		} else if nums[l]%2 == 1 {
			//右边有奇数, 移动左边找偶数
			l++
		}
	}

	return nums
}
