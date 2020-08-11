/**
 *
 * @author liangjf
 * @create on 2020/8/11
 * @version 1.0
 */
/**
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。



示例 1:

输入: [0,1,3]
输出: 2

输入: [0,1,2,3,4,5,6,7,9]
输出: 8


限制：

1 <= 数组长度 <= 10000
*/
package _3_2

//二分法 左右逼近
//通过mid下标来缩小区间, 最后定位到缺失的值, 或者到达末尾
func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		//下标为i的肯定是i, 如果不是, 证明前面有错位
		if nums[mid] > mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}
