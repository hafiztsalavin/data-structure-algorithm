func strStr(haystack string, needle string) int {
    
    sizeHaystack := len(haystack)
    sizeNeedle := len(needle)

    sizeCheck := sizeHaystack - sizeNeedle
    for i := 0; i <= sizeCheck; i++{
        candidate := haystack[i:i+sizeNeedle]
        if candidate == needle {
            return i
        }
        fmt.Println(candidate)
    }

    return -1
}