//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
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
//         /  \    / \
//        7    2  5   1
//
//
// 返回:
//
// [
//   [5,4,11,2],
//   [5,8,4,5]
//]
//
// Related Topics 树 深度优先搜索
// 👍 291 👎 0

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
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	one := make([]int, 0)
	_pathSum(root, sum, one, &res)
	return res
}

func _pathSum(root *TreeNode, sum int, one []int, res *[][]int) {
	if root == nil {
		return
	}

	one = append(one, root.Val)

	//到达叶子节点并且满足sum
	if sum == root.Val && root.Right == nil && root.Left == nil {
		//注意:深拷贝
		dst := make([]int, len(one))
		copy(dst, one)
		*res = append(*res, dst)

		return
	}

	_pathSum(root.Left, sum-root.Val, one, res)
	_pathSum(root.Right, sum-root.Val, one, res)
	//回溯
	one = one[:len(one)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
