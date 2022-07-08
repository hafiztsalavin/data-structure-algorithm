func strStr(haystack string, needle string) int {
    
    lengthHaystack := len(haystack)
    lengthNeedle := len(needle)
    
    if (lengthHaystack == 0 && lengthNeedle == 0) || haystack == needle {
        return 0
    } 
    
    if (lengthHaystack - lengthNeedle) < 0{
        return -1
    } 
    
    for i := 0; i < (lengthHaystack - lengthNeedle)+1; i++{
        if haystack[i:i+lengthNeedle] == needle{
            return i
        }
    }
    
    return -1
}