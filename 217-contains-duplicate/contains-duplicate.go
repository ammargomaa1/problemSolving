func containsDuplicate(nums []int) bool {
    mappedNums := map[int] bool{}

    for _, number := range  nums {
        _, ok := mappedNums[number]

        if !ok {
            mappedNums[number] = true
        } else {
            return true
        }
    }

    return false
}