/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dumbList := &ListNode{}
	current := dumbList

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				current.Next = &ListNode{
					Val: list2.Val,
				}
				list2 = list2.Next
				current = current.Next
			}else if list2.Val > list1.Val{
				current.Next = &ListNode{
					Val: list1.Val,
				}
				list1 = list1.Next
				current = current.Next
			} else if list1.Val == list2.Val {
				current.Next = &ListNode{
					Val: list1.Val,
					Next: &ListNode{
						Val: list2.Val,
					},
				}
				list1 = list1.Next
				list2 = list2.Next

				current = current.Next.Next
			}
		} else if list1 != nil && list2 == nil {
			current.Next = &ListNode{
				Val: list1.Val,
			}

			list1 = list1.Next
			current = current.Next
		}else {
			current.Next = &ListNode{
				Val: list2.Val,
			}

			list2 = list2.Next
			current = current.Next
		}
	}

	return dumbList.Next
}