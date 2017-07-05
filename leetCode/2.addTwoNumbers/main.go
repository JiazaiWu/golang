package main

type ListNode struct {
   Val int
   Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sum := 0;
    head := &ListNode{};
    node := &head;

    for {
        if l1 != nil {
            sum += l1.Val;
            l1 = l1.Next;
        }

        if l2 != nil {
            sum += l2.Val;
            l2 = l2.Next;
        }
        *node = &ListNode{Val: sum%10};
        sum /= 10;
        if l1 == nil && l2 == nil && sum == 0 {
            break;
        }
		node = &(*node).Next;
    }
    
    return head;
}
