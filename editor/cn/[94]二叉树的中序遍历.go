//给定一个二叉树，返回它的中序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表
// 👍 576 👎 0

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
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	dfs(root, &ret)
	return ret
}

func dfs(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, ret)
	*ret = append(*ret, node.Val)
	dfs(node.Right, ret)
}

//leetcode submit region end(Prohibit modification and deletion)
