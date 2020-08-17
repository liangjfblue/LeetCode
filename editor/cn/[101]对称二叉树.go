//给定一个二叉树，检查它是否是镜像对称的。
//
//
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//     1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//
//
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//     1
//   / \
//  2   2
//   \   \
//   3    3
//
//
//
//
// 进阶：
//
// 你可以运用递归和迭代两种方法解决这个问题吗？
// Related Topics 树 深度优先搜索 广度优先搜索
// 👍 957 👎 0

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	//return _dfs(root.Left, root.Right)
	return _bfs(root)
}

func _dfs(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}

	if n1.Val != n2.Val {
		return false
	}

	return _dfs(n1.Left, n2.Right) && _dfs(n1.Right, n2.Left)
}

func _bfs(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root, root}...)

	for len(queue) != 0 {
		n1, n2 := queue[0], queue[1]
		queue = queue[2:]

		if n1 == nil && n2 == nil {
			continue
		}

		if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		queue = append(queue, n1.Left)
		queue = append(queue, n2.Right)

		queue = append(queue, n1.Right)
		queue = append(queue, n2.Left)
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
