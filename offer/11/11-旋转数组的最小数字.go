/**
 *
 * @author liangjf
 * @create on 2020/7/25
 * @version 1.0
 */
/**
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,6,7,1,2]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0

输入：[1,2,5]
输出：1
*/
package _1

//遍历数组, 注意完全递增的情况, 这种情况直接就是第一个元素了
//时间复杂度O(n), 空间复杂度O(1)
func minArray(numbers []int) int {
	l := len(numbers)
	if l == 0 {
		return -1
	}

	for i := 0; i < l-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}

//分治法,左右两边同时扫描, 比遍历数组的方式缩短一半时间
//时间复杂度O(n), 空间复杂度O(1)
func minArray2(numbers []int) int {
	l := len(numbers)
	if l == 0 {
		return -1
	}

	l, r := 0, l-1
	for l < r {
		if numbers[l] > numbers[l+1] {
			return numbers[l+1]
		}

		if numbers[r-1] > numbers[r] {
			return numbers[r]
		}
		l++
		r--
	}
	return numbers[0]
}
