/**
 *
 * @author liangjf
 * @create on 2020/7/29
 * @version 1.0
 */
/**
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:

1 2 1 5 1 6 1 7 1
输入: [1, 2, 3, 2, 2, 2, 5, 4, 2]
输出: 2

限制：
1 <= 数组长度 <= 50000
*/
package _9

import "sort"

//排序取中位数: 时间O(nlogn)，空间O(1)
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	return nums[len(nums)/2]
}

//哈希法: 时间O(n)，空间O(n/2)
func majorityElement2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {
		m[nums[i]]++
		if m[nums[i]] > l/2 {
			return nums[i]
		}
	}

	return 0
}

//摩尔投票法
//大混战互相换命，不同的两者一旦遇见就同归于尽,最后活下来的值都是相同的，即要求的结果(计数器: 敌军自减, 友军自增)
//时间O(n)，空间O(1)
func majorityElement3(nums []int) int {
	res, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			//第一个人, 或者多一个人那方后面加回来了
			res = nums[i]
			count++
		} else {
			//第二个人开始
			if res == nums[i] {
				//自己人, 增加力量
				count++
			} else {
				//敌军, 互相拼了
				count--
			}
		}
	}
	return res
}
