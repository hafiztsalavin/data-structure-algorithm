func searchInsert(nums []int, target int) int {
    start := 0
    end := len(nums)-1
    result := binarySearch(nums, target, start, end)
    return result
}

func binarySearch(arr []int, target, start, end int )int{
    
    if start > end {
        return end + 1
    }
    
    mid := start + (end-start)/2
    
    if arr[mid] < target{
        return binarySearch(arr, target, mid+1, end)
    } else {
        return binarySearch(arr, target, start, mid-1)
    }
}
