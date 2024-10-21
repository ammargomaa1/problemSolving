func twoSum(nums []int, target int) []int {
    possibleResults := map[int]int{}
    for index, num := range nums {
        remaining := target - num
        
        if i, found := possibleResults[remaining]; found {
            return[]int{index, i}
        }

        possibleResults[num] = index
    }
    return []int{}
}