//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例：
//二叉树：[3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
//
// Related Topics 树 广度优先搜索
// 👍 598 👎 0

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		sum := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < sum; i++ {
			n := queue[0]
			queue = queue[1:]
			tmp = append(tmp, n.Val)

			if n.Left != nil {
				queue = append(queue, n.Left)
			}

			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		res = append(res, tmp)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
