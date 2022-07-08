func isPalindrome(x int) bool {
    inputNumber := x
    reverse := 0
    
    for inputNumber > 0 {
        pop := inputNumber % 10
        inputNumber = inputNumber/10
        
        reverse = (reverse * 10)+pop
    }
    
    if x == reverse {
        return true
    }
    return false
    
}