/**
 *
 * @author liangjf
 * @create on 2020/7/28
 * @version 1.0
 */
/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。


例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：

[3,9,20,15,7]


提示：

节点总数 <= 1000
*/
package _2_1

import "offer/model"

//层序遍历, 借助队列实现
func levelOrder(root *model.TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	queue := make([]*model.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		res = append(res, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}
