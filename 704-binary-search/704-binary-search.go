func search(nums []int, target int) int {
    size := len(nums)   
    if size == 1 {
        return singleArr(nums, target)
    } else if size == 2 {
        return doubleArr(nums, target)
    } else {
        
        return binarySearch(nums, 0, size-1, target)
    }
}

func singleArr(nums []int, target int) int {
    if nums[0] != target{
        return -1
    } else {
        return 0
    }
    
}

func doubleArr(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return i
        }
    }
    return -1
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
