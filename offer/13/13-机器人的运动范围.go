/**
 *
 * @author liangjf
 * @create on 2020/7/25
 * @version 1.0
 */
/**
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。
请问该机器人能够到达多少个格子？


示例 1：

输入：m = 2, n = 3, k = 1
输出：3
示例 2：

输入：m = 3, n = 1, k = 0
输出：1
提示：

1 <= n,m <= 100
0 <= k <= 20
*/
package _3

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, visited)
}

//深度优先遍历, 超某个方向去, 直至不符合条件回退, 然后再尝试其他方向
//+1是因为[0][0]肯定是符合的
func dfs(i, j, m, n, k int, visited [][]bool) int {
	//剪枝
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || (i/10+i%10+j/10+j%10) > k {
		return 0
	}

	//推进递归
	visited[i][j] = true
	return dfs(i+1, j, m, n, k, visited) +
		dfs(i-1, j, m, n, k, visited) +
		dfs(i, j+1, m, n, k, visited) +
		dfs(i, j-1, m, n, k, visited) + 1
}
