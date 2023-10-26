func maxPower(s string) int {
    globalMax := 1
    currMax := 1
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1]{
            currMax += 1
        } else{
            if currMax > globalMax {
                globalMax = currMax
            }
            currMax = 1
        }

        if currMax > globalMax {
                globalMax = currMax
            }
    }
    return globalMax
}