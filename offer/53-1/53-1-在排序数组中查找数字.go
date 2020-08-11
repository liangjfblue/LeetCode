/**
 *
 * @author liangjf
 * @create on 2020/8/10
 * @version 1.0
 */
/**
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


限制：

0 <= 数组长度 <= 50000
*/
package _3_1

//遍历
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	hadFound := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			sum++
			hadFound = true
		} else {
			if hadFound == true {
				break
			}
		}
	}

	return sum
}

//二分法
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			//往前后找相等的
			tmp := mid
			for mid >= 0 && nums[mid] == target {
				mid--
				sum++
			}

			mid = tmp + 1
			for mid < len(nums) && nums[mid] == target {
				mid++
				sum++
			}
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}

	return sum
}
