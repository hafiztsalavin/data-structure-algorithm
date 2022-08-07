func intersection(nums1 []int, nums2 []int) []int {
    result := []int{}
    
    hashNums := map[int]int{}
    for _, num := range nums1 {        
        hashNums[num] = 1
    }
    
    for i := 0; i < len(nums2); i++{
        if value, key := hashNums[nums2[i]]; key{
            if value != 0 {
                result = append(result, nums2[i])
                hashNums[nums2[i]] = 0            
            }
        }
    }
    
    return result
}