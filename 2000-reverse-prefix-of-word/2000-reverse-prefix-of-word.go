func reversePrefix(word string, ch byte) string {
   for i := 0; i < len(word); i++ {
       if word[i] == ch {
           reverse := reverseString(word[:i+1])
           result := reverse + word[i+1:]
           return result
 
       }
   } 
   return word
}

func reverseString(words string) (result string) {
    for _, word := range words {
        result = string(word) + result
    }
    return
}