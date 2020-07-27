/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。


示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


限制：

0 <= matrix.length <= 100
0 <= matrix[i].length <= 100
*/
package _9

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return nil
	}

	res := make([]int, 0)

	_bfs(matrix, 0, 0, &res)
	return res
}

func _bfs(matrix [][]int, x, y int, res *[]int) {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || matrix[x][y] == '*' {
		return
	}

	*res = append(*res, matrix[x][y])

	matrix[x][y] = '*'
	_bfs(matrix, x, y+1, res)
	_bfs(matrix, x+1, y, res)
	_bfs(matrix, x, y-1, res)
	_bfs(matrix, x-1, y, res)

	return
}

func spiralOrder2(matrix [][]int) []int {
	direction := 0
	index := 0
	l := len(matrix) * len(matrix[0])
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	res := make([]int, 0)

	for index < l {
		if direction%4 == 0 {
			// 左 -> 右
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		} else if direction%4 == 1 {
			// 上 -> 下
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		} else if direction%4 == 2 {
			// 右 -> 左
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		} else {
			// 下 -> 上
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
		direction++
	}

	return res
}
