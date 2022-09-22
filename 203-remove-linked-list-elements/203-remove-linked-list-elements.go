/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    // handle if head = nil
    if head == nil{
        return head
    }
    
    last := head            
    for last.Next != nil{
        // handle value if in the mid and last
        if last.Next.Val == val{
            last.Next = last.Next.Next
        } else {
            last = last.Next
        }
    }
    
    // handle if value is head
    if head.Val == val{
        return head.Next
    }
    
    return head
}