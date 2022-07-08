func twoSum(nums []int, target int) []int {
    numsMap := map[int]int{}
    
    for i, num := range nums {
        rest := target - num
        
        if value, key := numsMap[rest]; key{
            return []int{value, i}
        }
        
        numsMap[num] = i
    }
    return []int{}
}