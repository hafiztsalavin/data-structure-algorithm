func sortArray(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
        
    left := sortArray(nums[:len(nums)/2])
    right := sortArray(nums[len(nums)/2:])
    
    return merge(left, right)
}

func merge(first, second []int) []int {
    size := len(first) + len(second)
    result := make([]int, size)
    
    for i, j, k := 0, 0, 0; k < size; k++{
        if i > len(first)-1 && j <= len(second)-1 {
            result[k] = second[j]
            
            j++
        } else if j > len(second)-1 && i <= len(first)-1 {
            result[k] = first[i]
            
            i++
        } else if first[i] < second[j] {
            result[k] = first[i]
        
            i++
        } else {
            result[k] = second[j]
            
            j++
        }
    }
    
    return result
}
