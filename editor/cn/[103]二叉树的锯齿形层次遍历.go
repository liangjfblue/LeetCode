//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
// 例如：
//给定二叉树 [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回锯齿形层次遍历如下：
//
// [
//  [3],
//  [20,9],
//  [15,7]
//]
//
// Related Topics 栈 树 广度优先搜索
// 👍 250 👎 0

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	flag := true
	for len(queue) != 0 {
		sum := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < sum; i++ {
			n := queue[0]
			queue = queue[1:]

			//每层头尾插入变换 先尾插, 后头插
			if flag {
				//尾插
				tmp = append(tmp, n.Val)
			} else {
				//头插
				tmp = append([]int{n.Val}, tmp...)
			}

			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}

		flag = !flag
		res = append(res, tmp)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
