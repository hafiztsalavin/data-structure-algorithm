func maxProfit(prices []int) int {
    maxSell, min := 0, prices[0]
    for i := 1; i < len(prices); i++ {
        if price < min {
            min = price
        } else if price - min > maxSell {
            maxSell = price
        }
    }
   return maxSell
}