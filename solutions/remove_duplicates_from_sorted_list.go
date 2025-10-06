package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	pointer := head

	for pointer != nil && pointer.Next != nil {
		for pointer.Next != nil && pointer.Next.Val == pointer.Val {
			nextPointer := pointer.Next
			pointer.Next = pointer.Next.Next
			nextPointer.Next = nil
		}
		pointer = pointer.Next
	}

	return head
}
