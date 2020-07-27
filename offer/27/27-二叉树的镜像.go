/**
 *
 * @author liangjf
 * @create on 2020/7/27
 * @version 1.0
 */
package _7

import "offer/model"

//层序遍历, 交换左右子树就行
func mirrorTree(root *model.TreeNode) *model.TreeNode {
	queue := make([]*model.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		//pop queue
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if node == nil {
			continue
		}

		//swap left right
		node.Left, node.Right = node.Right, node.Left

		//push queue
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

//递归法, 前序遍历, 并且调换左右子树
func mirrorTree2(root *model.TreeNode) *model.TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	mirrorTree2(root.Left)
	mirrorTree2(root.Right)

	return root
}
