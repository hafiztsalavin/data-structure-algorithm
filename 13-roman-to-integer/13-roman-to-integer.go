func romanToInt(s string) int {
    roman := map[string]int {
        "I" : 1,
        "V" : 5,
        "X" : 10,
        "L" : 50,
        "C" : 100,
        "D" : 500,
        "M" : 1000,
    }
    
    total := 0
    size := len(s)
    
    for i := 0; i < size; i++{
        if i  == size-1{
            total += roman[string(s[i])]
        }
        
        if i + 1 < size {
            if roman[string(s[i])] < roman[string(s[i+1])] {
                total -= roman[string(s[i])]
            } else {
                total += roman[string(s[i])]
            }
        }
    }
    
    return total
    
}