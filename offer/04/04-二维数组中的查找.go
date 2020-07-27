/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。


限制：

0 <= n <= 1000

0 <= m <= 1000
*/
package _4

/**
从左下角开始走，利用这个顺序关系可以在O(m+n)的复杂度下解决这个题：
如果当前位置元素比target小，则行row--
如果当前位置元素比target大，则列col++
如果相等，返回true
如果越界了还没找到，说明不存在，返回false
*/
//时间复杂度O(logN), 空间复杂度O(1)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[0]) {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] {
			i--
		}
	}

	return false
}
