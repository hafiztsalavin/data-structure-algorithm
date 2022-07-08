func plusOne(digits []int) []int {
    size := len(digits)-1
    
    ans := divideIncrement(digits, size)
    return ans
}


func divideIncrement( arr []int, size int ) []int{
    
    if arr[size] != 9{
        arr[size] = arr[size]+1
    } else if size != 0{
        arr[size] = 0
        return divideIncrement(arr, size-1)
    } else {
        res := make([]int, len(arr)+1)
        res [0] = 1
        return res
    }
    return arr
}
