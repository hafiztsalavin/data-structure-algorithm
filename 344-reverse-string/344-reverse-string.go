func reverseString(s []byte)  {
    pointer1 := 0
    pointer2 := len(s)-1
    
    for pointer1 < pointer2 {
        s[pointer1], s[pointer2] = s[pointer2], s[pointer1]
    
        pointer1 += 1
        pointer2 -= 1
    }
    
}