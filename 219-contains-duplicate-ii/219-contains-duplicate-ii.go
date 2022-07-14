func containsNearbyDuplicate(nums []int, k int) bool {
    hashMap := map[int]int{}
    for i := 0; i < len(nums); i++{
        value, key := hashMap[nums[i]]
        if key && (check(value, i)<=k){
            return true
        }
        hashMap[nums[i]] = i    
    }
    
    return false
}

func check(a, b int) int {
    if (a-b) < 0{
        return -(a-b)
    }
    return (a-b)
}