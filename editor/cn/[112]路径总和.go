//给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例:
//给定如下二叉树，以及目标和 sum = 22，
//
//               5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \      \
//        7    2      1
//
//
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
// Related Topics 树 深度优先搜索
// 👍 405 👎 0

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
//func hasPathSum(root *TreeNode, sum int) bool {
//	res := false
//	_hasPathSum(root, sum, &res)
//	return res
//}
//
//func _hasPathSum(root *TreeNode, sum int, res *bool) {
//	if root == nil {
//		return
//	}
//
//	//到达叶子节点并且满足sum
//	if sum == root.Val && root.Right == nil && root.Left == nil {
//		*res = true
//		return
//	}
//
//	sum -= root.Val
//	_hasPathSum(root.Left, sum, res)
//	_hasPathSum(root.Right, sum, res)
//	//回溯
//	sum += root.Val
//
//	return
//}
func hasPathSum(root *TreeNode, sum int) bool {

	return _hasPathSum(root, sum)
}

func _hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	//到达叶子节点并且满足sum
	if sum == root.Val && root.Right == nil && root.Left == nil {
		return true
	}

	return _hasPathSum(root.Left, sum-root.Val) || _hasPathSum(root.Right, sum-root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
