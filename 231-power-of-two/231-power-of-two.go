func isPowerOfTwo(n int) bool {
    if n == 1{
        return true
    }
    
    if n % 2 != 0 || n == 0 || n < 1{
        return false
    } 
    
    
    nums := n
    for nums > 1 {
        if nums % 2 == 1 {
            return false
        }
        
        nums /= 2
    }
    
    return true
}