//给定一个二叉树，判断它是否是高度平衡的二叉树。
//
// 本题中，一棵高度平衡二叉树定义为：
//
//
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
//
//
// 示例 1:
//
// 给定二叉树 [3,9,20,null,null,15,7]
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回 true 。
//
//示例 2:
//
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//        1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
//
//
// 返回 false 。
//
//
// Related Topics 树 深度优先搜索
// 👍 423 👎 0

package cn

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return _isBalanced(root) >= 0
}

func _isBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := _isBalanced(root.Left)
	rightHeight := _isBalanced(root.Right)
	if leftHeight == -1 || rightHeight == -1 || math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return -1
	}

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

//leetcode submit region end(Prohibit modification and deletion)
