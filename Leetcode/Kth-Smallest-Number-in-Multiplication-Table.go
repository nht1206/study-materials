package leetcode

// ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head, Val: -999999}
	cur := newHead
	last := newHead
	front := head
	for front.Next != nil {
		if front.Val == cur.Val {
			// fmt.Printf("Same node front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
			front = front.Next
			continue
		} else {
			if cur.Next != front {
				// fmt.Printf("Delete duplicate nodes front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
				last.Next = front
				if front.Next != nil && front.Next.Val != front.Val {
					last = front
				}
				cur = front
				front = front.Next
			} else {
				// fmt.Printf("front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
				last = cur
				cur = cur.Next
				front = front.Next
				// fmt.Printf("front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val) after regular loop

			}
		}
	}
	if front.Val == cur.Val {
		// fmt.Printf("Same node front = %v | cur = %v | last = %v\n", front.Val, cur.Val, last.Val)
		last.Next = nil
	} else {
		if cur.Next != front {
			last.Next = front
		}
	}
	return newHead.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next != nil && head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates(head.Next)
	}
	head.Next = deleteDuplicates(head.Next)
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// Simple solution for double loop O(n*m)
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	head = nilNode

	lastVal := 0
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			lastVal = head.Next.Val
			for head.Next != nil && lastVal == head.Next.Val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return nilNode.Next
}

// Double pointer + delete flag, single loop solution O(n)
func deleteDuplicates4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	// The last traversal has the flag bit of the delete operation
	lastIsDel := false
	// virtual empty node
	head = nilNode
	// Front and back pointers are used to judge
	pre, back := head.Next, head.Next.Next
	// Only delete the previous duplicate element each time, leaving one for the next traversal judgment
	// The update position and value of the pre and back pointers are more important and ingenious
	for head.Next != nil && head.Next.Next != nil {
		if pre.Val != back.Val && lastIsDel {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
			continue
		}

		if pre.Val == back.Val {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = true
		} else {
			head = head.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
		}
	}
	// handle the case of [1,1] when there is one left to delete
	if lastIsDel && head.Next != nil {
		head.Next = nil
	}
	return nilNode.Next
}
