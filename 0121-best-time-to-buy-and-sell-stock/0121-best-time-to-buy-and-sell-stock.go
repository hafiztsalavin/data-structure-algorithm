func maxProfit(prices []int) int {

    maxSell, min := 0, prices[0]

    for _, price := range prices{
        if price < min{
            min = price
        } else if price - min > maxSell {
            maxSell = price - min
        }
    }
    return maxSell
}