func findRestaurant(list1 []string, list2 []string) []string {
    result := []string{}
    max := len(list1) + len(list2)
    
    mapRestaurant := map[string]int{}
    for i, restaurant := range list1 {
        mapRestaurant[restaurant] = i+1 
    }

    for i, restaurant2 := range list2{
        if value, restaurant := mapRestaurant[restaurant2]; restaurant{
            tempMax := (i+1)+value
            
            if tempMax < max {
                max = tempMax
                result = result[:0]
                result = append(result, restaurant2)
            } else if tempMax == max{
                result = append(result, restaurant2)
            }
        }
    }
    return result    
}
    