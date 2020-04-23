package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next:nil}
	cur :=head
	carry:= 0
	for l1 != nil || l2 != nil {
		var x,y int
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}
		sum := x+y+carry
		cur.Next = &ListNode{Val:sum%10,Next: nil}
		cur = cur.Next
		carry = sum / 10
		if l1 !=nil {
			l1 = l1.Next
		}
		if l2 !=nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry,Next: nil}
	}
	return head.Next
}
