func twoSum(nums []int, target int) []int {
    for index, num1 := range nums {

        for i := index + 1; i < len(nums); i++ {
            if i < len(nums) && num1 + nums[i] == target {
                return []int{index, i}
            }
        }
    }
    return []int{}
}