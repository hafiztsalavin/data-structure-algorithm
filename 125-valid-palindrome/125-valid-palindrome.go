func isPalindrome(s string) bool {

    size := len(s)
    i := 0
    j := size-1
    for i < j {
        right := check(s[i])
        left := check(s[j])
        
        if right == 0{
            i++
        }
        
        if left == 0{
            j--
        }
        
        if right == 0 || left == 0{
            continue
        }
        
        if right != left{
            return false
        }
        
        i++
        j--
    }
  
    
    return true
    
}

func check(b byte)byte{
    if (b >= 97 && b <= 122) || (b >= 48 && b <= 57 ){
        return b
    } else if (b >= 65 && b <= 90)  {
        b += 32
        return b
    }
    
    return 0
}