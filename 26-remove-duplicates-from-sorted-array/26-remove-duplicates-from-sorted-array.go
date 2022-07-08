func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    curr := nums[0]
    fill := 1
    
    
    for i := 1; i < len(nums);i++{
        if nums[i] != curr{
            
            curr = nums[i]
            nums[fill]= curr
            fill++
        }
    }
    
    return fill
}