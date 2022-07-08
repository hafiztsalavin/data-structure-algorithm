/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    node := ListNode{}
    newList := &node
    
    for list1 != nil && list2 != nil {
        if list1.Val > list2.Val {          
            newList.Next = list2
            list2 = list2.Next
        } else {
            newList.Next = list1
            list1 = list1.Next
        }
        
        newList = newList.Next
    }
    
    // sisa nya jika salah satu list abis
    if list1 != nil {
        newList.Next = list1
    }
    
    if list2 != nil {
        newList.Next = list2
    }
    
    return node.Next
}
