func sortColors(nums []int) []int {
    
    size := len(nums)
    for i := 0; i < size; i++{
        minIndexValue := i
        for j := i+1; j < size; j++{
            if nums[j] < nums[minIndexValue] {
                minIndexValue = j
            }
        }
        nums[i], nums[minIndexValue] = nums[minIndexValue], nums[i]
    }
    
    return nums
}