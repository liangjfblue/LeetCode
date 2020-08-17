//给定一个二叉树，找出其最小深度。
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//
// 给定二叉树 [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最小深度 2.
// Related Topics 树 深度优先搜索 广度优先搜索
// 👍 319 👎 0

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//到达叶子节点, 直接返回1
	if root.Left == nil && root.Right == nil {
		return 1
	}

	//计算左右子树的深度
	l := minDepth(root.Left)
	r := minDepth(root.Right)

	//其中一个为nil,证明其到达了叶子节点, 可以返回了
	if root.Left == nil || root.Right == nil {
		return l + r + 1
	}

	//if root.Left == nil {
	//	return l + 1
	//}
	//
	//if root.Right == nil {
	//	return r + 1
	//}

	//都不为nil, 取较小的+1
	return int(math.Min(float64(l), float64(r))) + 1
}

//leetcode submit region end(Prohibit modification and deletion)
