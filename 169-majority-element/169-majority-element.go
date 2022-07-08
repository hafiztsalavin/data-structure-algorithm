func majorityElement(nums []int) int {
    hashArr := make(map[int]int)
    
    for _, num := range nums {
        hashArr[num] += 1
    }
   
    count := 0
    max := 0
    
    for key, value := range hashArr {
        if value > count {
            count = value
            max = key
        }
    }
    
    return max
}
