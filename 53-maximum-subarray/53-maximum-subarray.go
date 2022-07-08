func maxSubArray(nums []int) int {
    size := len(nums)-1
    ans := findSubArray(nums, 0, size)
    return ans
}

func findSubArray(arr []int, first, last int) int {
    if first == last {
        return arr[first]
    } else {
        mid := (first+last)/2
        
        maxLeft := findSubArray(arr, first, mid)
        maxRight := findSubArray(arr, mid+1, last)
        between := findBetween(arr, first, last, mid)
        
        return max(maxLeft, maxRight, between)
    }
}

func findBetween(arr []int, first, right, mid int) int{
    sumRight := -1<<31
    sumLeft := -1<<31
    
    sum := 0
    for i := mid; i >= first; i--{
        sum += arr[i]
        
        sumLeft = max(sum, sumLeft)
    }
    
    sum = 0
    for i := mid+1; i <= right; i++{
        sum += arr[i]
        sumRight = max(sum, sumRight)
    }
    
    return sumLeft+sumRight
}

func max( values ...int) int {
    max := values[0]
    
    for _, value := range values{
        if value > max{
            max = value
        }
    }
    return max
}
