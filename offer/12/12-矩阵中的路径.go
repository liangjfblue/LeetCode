/**
 *
 * @author liangjf
 * @create on 2020/7/25
 * @version 1.0
 */
/**
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，
每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。


示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
提示：

1 <= board.length <= 200
1 <= board[i].length <= 200
*/
package _2

//思路比较简单, 遍历二维数组, 找到和word首字符相等的, 然后从当前位置开始上下左右探索是否和word的字符一致
func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 || word == "" {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(board, i, j, word, 0) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j int, word string, index int) bool {
	//index是word的下标, 若index和word长度相等, 说明已经找到word的路径
	if index == len(word) {
		return true
	}

	//上下左右边界, 字符相等判断
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] != word[index] {
		return false
	}

	//先标明已遍历
	board[i][j] = '*'

	//上下左右探索, 利用了深度优先遍历的思想
	ret := dfs(board, i+1, j, word, index+1) ||
		dfs(board, i-1, j, word, index+1) ||
		dfs(board, i, j+1, word, index+1) ||
		dfs(board, i, j-1, word, index+1)

	//到这里, 证明index下标和word字符相等, 需要还原给后面的判断用
	board[i][j] = word[index]

	return ret
}
