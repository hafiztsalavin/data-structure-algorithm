/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    firstPointerNode := head
    secondPointerNode := head
    
    for secondPointerNode != nil && secondPointerNode.Next != nil{
        firstPointerNode = firstPointerNode.Next
        secondPointerNode = secondPointerNode.Next.Next
    }
    
    return firstPointerNode
}