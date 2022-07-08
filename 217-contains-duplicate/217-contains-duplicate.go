func containsDuplicate(nums []int) bool {
    
    // key = num, value = index
    duplicatesMap := make(map[int]int)
    
    size := len(nums)
    for i := 0; i < size; i++{
        if _, value := duplicatesMap[nums[i]]; value {
            return true
        }
        
        duplicatesMap[nums[i]] = i
    }
    fmt.Println(duplicatesMap)  
    
    
    return false
    
}