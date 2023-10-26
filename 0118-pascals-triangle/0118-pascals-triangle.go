func generate(numRows int) [][]int {
    ans := [][]int{{1}}

    for i := 1; i < numRows; i++{
        prevArr := ans[i-1] 

        newArr := make([]int,i+1)
        newArr[0], newArr[i] = 1,1
        if i > 1 {
            for j:= 0; j < i-1; j++{
                num := prevArr[j] + prevArr[j+1]
                newArr[j+1] = num
            }

        }   
        ans = append(ans, newArr)
    }

    return ans
    
}