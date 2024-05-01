/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dumbList := &ListNode{}
	current := dumbList

	curry  := 0
	for l1 != nil || l2 != nil || curry != 0   {
		sum := curry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}



		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		curry = sum / 10
		
	}

	return dumbList.Next
}