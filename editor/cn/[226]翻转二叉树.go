//翻转一棵二叉树。
//
// 示例：
//
// 输入：
//
//      4
//   /   \
//  2     7
// / \   / \
//1   3 6   9
//
// 输出：
//
//      4
//   /   \
//  7     2
// / \   / \
//9   6 3   1
//
// 备注:
//这个问题是受到 Max Howell 的 原问题 启发的 ：
//
// 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
// Related Topics 树
// 👍 536 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	return dfs222(root)
}

func dfs222(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		n.Left, n.Right = n.Right, n.Left

		if n.Left != nil {
			queue = append(queue, n.Left)
		}

		if n.Right != nil {
			queue = append(queue, n.Right)
		}
	}

	return root
}

func _invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	right := _invertTree(root.Right)
	left := _invertTree(root.Left)

	root.Right = left
	root.Left = right

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
