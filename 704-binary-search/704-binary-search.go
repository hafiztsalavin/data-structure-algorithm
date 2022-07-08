func search(nums []int, target int) int {
    size := len(nums)   
   
    return binarySearch(nums, 0, size-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
    if end < start || start > end {
		return -1
	}
    
    mid := start + (end-start)/2
    
    if nums[mid] == target{
        return mid
    } else if nums[mid] < target{
        return binarySearch(nums, mid+1, end, target)
    } else {
        return binarySearch(nums, start, mid-1, target)
    }
}
