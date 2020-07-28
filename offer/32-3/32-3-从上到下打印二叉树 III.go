/**
 *
 * @author liangjf
 * @create on 2020/7/28
 * @version 1.0
 */
/**
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。


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
  [20,9],
  [15,7]
]



提示：

节点总数 <= 1000
*/
package _2_3

import "offer/model"

//每层遍历的都会加进去队列, 所以遍历拿出队列当前的大小,就是当前层的节点个数,
//然后再通过循环遍历每个节点再把他们的左右子节点再进去队列, 下次循环,
//这样就实现层序遍历和每层换一行了
//注意: 通过flag来切换 切片的插入方式, 头部插入还是尾部, 实现正反序
func levelOrder(root *model.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*model.TreeNode, 0)
	queue = append(queue, root)

	flag := true
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

			//通过flag来切换 切片的插入方式, 头部插入还是尾部, 实现正反序
			if flag {
				one = append(one, node.Val)
			} else {
				one = append([]int{node.Val}, one...)
			}
		}

		flag = !flag
		res = append(res, one)
	}
	return res
}
