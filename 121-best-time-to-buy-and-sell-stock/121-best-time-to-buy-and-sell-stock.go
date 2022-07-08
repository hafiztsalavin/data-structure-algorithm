func maxProfit(prices []int) int {
    maxSell := 0
    min := prices[0]
    
    for _, price := range prices{
        if price < min {
            min = price
        } else if price - min > maxSell{
            maxSell = price - min
        }
    }
    return maxSell
}