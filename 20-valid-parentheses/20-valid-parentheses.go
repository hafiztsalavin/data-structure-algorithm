func isValid(s string) bool {
    stackBracket := []string{}
    bracketMap := map[string]string{"{" : "}", "(":")", "[":"]"}
    
    if len(s)%2 == 1{
        return false
    }
    
    for _, value := range s{
        
        if _, open := bracketMap[string(value)]; open{
            stackBracket = append(stackBracket, string(value))
        } else {
            if len(stackBracket) != 0{
                size := len(stackBracket)-1
                top := stackBracket[size]
                    
                if bracketMap[top] != string(value){
                    return false
                }
                
                stackBracket = stackBracket[:size]
            } else {
                return false
            }
        }
    }
    
    if len(stackBracket) != 0{
        return false
    } else {
        return true
    }
    
}