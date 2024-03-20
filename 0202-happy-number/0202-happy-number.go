func isHappy(n int) bool {
    listMeet := map[int]bool{}
    for {
        numNext := 0
        for ; n > 0; n = n / 10{
            numNext += (n%10)*(n%10)
        }

        if _, ok := listMeet[numNext]; ok{
            return numNext == 1
        } else {
            listMeet[numNext] = true
        }

        n = numNext
    }

    return false  
}