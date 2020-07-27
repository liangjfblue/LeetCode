/**
 *
 * @author liangjf
 * @create on 2020/7/24
 * @version 1.0
 */
/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

限制：

0 <= 链表长度 <= 10000
*/
package _6

import "offer/model"

//顺序遍历, 翻转数组
//时间复杂度O(n), 空间复杂度O(n)
func reversePrint(head *model.ListNode) []int {
	if head == nil {
		return nil
	}

	p := head
	ret := make([]int, 0)
	for p != nil {
		ret = append(ret, p.Val)
		p = p.Next
	}

	l := len(ret)
	for i := 0; i < l/2; i++ {
		ret[i], ret[l-1-i] = ret[l-1-i], ret[i]
	}

	return ret
}

//利用递归, 实现从后到前存储val
//时间复杂度O(n), 空间复杂度O(n)
func reversePrint2(head *model.ListNode) []int {
	if head == nil {
		return nil
	}

	ret := reversePrint2(head.Next)
	ret = append(ret, head.Val)

	return ret
}
