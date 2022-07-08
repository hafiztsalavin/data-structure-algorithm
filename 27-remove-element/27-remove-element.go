func removeElement(nums []int, val int) int {
    start := 0
    end := len(nums)-1
    
    for start <= end {
        if nums[start] != val {
            start++
        } else if nums[end] == val {
            end--
        } else {
            nums[start], nums[end] = nums[end], nums[start]
            start++
            end--
        }
    }
    return start
}