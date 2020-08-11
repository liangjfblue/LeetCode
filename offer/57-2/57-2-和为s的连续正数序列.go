/**
 *
 * @author liangjf
 * @create on 2020/8/11
 * @version 1.0
 */
/**
输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。



示例 1：

输入：target = 9
输出：[[2,3,4],[4,5]]
示例 2：

输入：target = 15
输出：[[1,2,3,4,5],[4,5,6],[7,8]]


限制：

1 <= target <= 10^5
*/
package _7_2

func findContinuousSequence(target int) [][]int {
	if target == 1 {
		return [][]int{{1}}
	}

	res := make([][]int, 0)
	l, r := 1, 2
	sum := l + r
	for l < r {
		if sum == target {
			tmp := make([]int, 0)
			for i := l; i <= r; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			sum -= l
			l++
		} else if sum > target {
			sum -= l
			l++
		} else {
			r++
			sum += r
		}
	}
	return res
}
