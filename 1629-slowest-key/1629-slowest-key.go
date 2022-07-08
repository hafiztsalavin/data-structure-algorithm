func slowestKey(releaseTimes []int, keysPressed string) byte {
    
    maxTimes := releaseTimes[0]-0
    maxIndex := 0
    
    for i:= 1; i < len(releaseTimes); i++{
        tempTimes := releaseTimes[i]- releaseTimes[i-1]
        if (tempTimes > maxTimes) || (tempTimes == maxTimes && keysPressed[i] > keysPressed[maxIndex]){
            maxTimes = tempTimes
            maxIndex = i
        }
    }
    
    return keysPressed[maxIndex]
}