func reverse(x int) int {
    maxInt := 1 << 31
    minInt := -1 << 31
    
    num := x
    reverseInt := 0
    pop := 0  
    
    for num != 0 {
        pop = num % 10
        num = num/10
        reverseInt  = (reverseInt * 10) + pop
    }  
    
    if reverseInt > maxInt || reverseInt < minInt {
        return 0
    }
    
    return reverseInt
}