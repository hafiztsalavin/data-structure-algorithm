func findDuplicate(nums []int) int {
    // formula for expectation is (n * (n + 1)/2)
    
    dup := map[int]int{}
    for _, num := range nums{
        if _, key := dup[num]; key{
         return num   
        }
        dup[num] = 1
    }
    
    return 0
}