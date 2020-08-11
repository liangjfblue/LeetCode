/**
 *
 * @author liangjf
 * @create on 2020/8/10
 * @version 1.0
 */
package _2

import "offer/model"

//相交   a+same+b = b+same+a
//不相交 a+b = b+a (a=nil b=nil)
func getIntersectionNode(headA, headB *model.ListNode) *model.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}
