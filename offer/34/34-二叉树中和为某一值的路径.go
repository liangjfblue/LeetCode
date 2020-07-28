/**
 *
 * @author liangjf
 * @create on 2020/7/28
 * @version 1.0
 */
/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。


示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]


提示：

节点总数 <= 10000
*/
package _4

import "offer/model"

/**
思路:
1.经典的回溯问题
2.通过深度遍历, 探索各条路径
3.若符合sum就返回
4.若不满足, 减掉list末端的节点, 继续尝试左右子树的探索
5.这里一个坑是go中切片的底层数据结构, 在append到结果时需要通过深拷贝, 不然叶子节点会在回溯时被删掉, 就不正确了
	tmp := make([]int, len(one))
	copy(tmp, one)
*/

//两次遍历, 第一次结束条件是nil, 第二次是和为sum或者nil
func pathSum(root *model.TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	one := make([]int, 0)
	bfs(root, &res, sum, one)
	return res
}

func bfs(root *model.TreeNode, res *[][]int, sum int, one []int) {
	if root == nil {
		return
	}

	sum -= root.Val
	one = append(one, root.Val)
	if sum == 0 && root.Right == nil && root.Left == nil {
		//坑: 防止下面回溯时剪掉叶子节点
		tmp := make([]int, len(one))
		copy(tmp, one)

		*res = append(*res, tmp)
	}

	bfs(root.Left, res, sum, one)
	bfs(root.Right, res, sum, one)
	//回溯删掉节点, 并且加回sum减掉的值
	sum += root.Val
	one = one[:len(one)-1]
}
