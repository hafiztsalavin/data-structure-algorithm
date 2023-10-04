func romanToInt(s string) int {

    romanDict := map[string]int{
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
        if i == size -1 {
            total += romanDict[string(s[i])]
        }

        if i+1 < size{
            if romanDict[string(s[i])] < romanDict[string(s[i+1])] {
                total -= romanDict[string(s[i])]
            } else {
                total += romanDict[string(s[i])]
            }
        }
    }

    return total
}