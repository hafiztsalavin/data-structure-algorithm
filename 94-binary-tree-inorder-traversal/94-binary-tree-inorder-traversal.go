/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
 	result := []int{}

	curr := root

	for curr != nil {
		if curr.Left == nil {
			result = append(result, curr.Val)
			curr = curr.Right
		} else {
			pre := curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}

			pre.Right = curr
			temp := curr
			curr = curr.Left
			temp.Left = nil
		}
	}
	return result
    
}