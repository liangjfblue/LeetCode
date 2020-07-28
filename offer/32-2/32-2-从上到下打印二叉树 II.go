/**
 *
 * @author liangjf
 * @create on 2020/7/28
 * @version 1.0
 */
/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。


例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]


提示：

节点总数 <= 1000
*/
package _2_2

import "offer/model"

//每层遍历的都会加进去队列, 所以遍历拿出队列当前的大小,就是当前层的节点个数,
//然后再通过循环遍历每个节点再把他们的左右子节点再进去队列, 下次循环,
//这样就实现层序遍历和每层换一行了
func levelOrder(root *model.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*model.TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		one := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			one = append(one, node.Val)
		}

		res = append(res, one)
	}
	return res
}
