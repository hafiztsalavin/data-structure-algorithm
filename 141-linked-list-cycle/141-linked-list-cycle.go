/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {

    first := head
    second := head
    
    for (second != nil && second.Next != nil){
        second = second.Next.Next
        first = first.Next
        if first == second{
            return true
        }
    } 
    
    return false
}