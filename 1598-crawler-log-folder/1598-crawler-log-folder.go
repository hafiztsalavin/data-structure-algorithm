func minOperations(logs []string) int {
    stackFolder := []string{}
    deleteFolder := "../"
    stayFolder := "./"
    
    for _, log := range logs{

        sizeFolder := len(stackFolder)
        if sizeFolder != 0{
            
            if log == deleteFolder{
                stackFolder = stackFolder[:sizeFolder-1]
            } else if log == stayFolder{
                continue
            } else {
                stackFolder = append(stackFolder, log)
            }
        } else {
            if log == stayFolder || log == deleteFolder{
                continue
            } else {
                stackFolder = append(stackFolder, log)
                
            }   
        }
        
    }

    fmt.Println(stackFolder)
    return len(stackFolder)
}