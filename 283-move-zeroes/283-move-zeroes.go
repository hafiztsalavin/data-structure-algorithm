func moveZeroes(nums []int)  {
    size := len(nums)
    
    firstPointer := 0
    secondPointer := 0
    j := size
    for firstPointer < j{
        if nums[firstPointer] != 0 {
            nums[firstPointer], nums[secondPointer] = nums[secondPointer], nums[firstPointer]
            firstPointer++
            secondPointer++
        } else {
            firstPointer++
        }
    }
}