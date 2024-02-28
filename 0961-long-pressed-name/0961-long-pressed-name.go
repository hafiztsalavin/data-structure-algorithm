func isLongPressedName(name string, typed string) bool {
    countCharName, nameLen := 0, len(name)
    var prevCharName byte
    for i := 0; i < len(typed); i++ {
        if countCharName < nameLen && name[countCharName] == typed[i]{
            countCharName ++
            prevCharName = typed[i]
        } else if prevCharName != typed[i]{
            return false
        }
    }
    return countCharName == nameLen
}

