func thirdMax(nums []int) int {
    min := -1 << 63
    first := min
    second := min
    third := min
    
    for _, value := range nums{
        if value == first || value == second || value == first{
            continue
        } 
        
        if value > first {
            third = second
            second = first
            first = value
        } else if value > second {
            third = second
            second = value
        } else if value > third {
            third = value
        }
    }
    
    if second == min || third == min {
        return first
    }
    
    return third
}
