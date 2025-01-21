func productExceptSelf(nums []int) []int {
    totalProduct := 1
    outPut := make([]int, len(nums))
    numberOfZeros := 0
    indexOfZero := 0
    indexOfZeroChanged := false
    for index, num := range nums {
        if num == 0 {
            numberOfZeros++
            indexOfZero = index
            indexOfZeroChanged = true
            continue
        }

        totalProduct = totalProduct * num
    }


    if numberOfZeros > 1 {
        return outPut
    }

    if indexOfZeroChanged {
        outPut[indexOfZero] = totalProduct
        return outPut
    }

    for currentIndex, num := range nums {
        
        outPut[currentIndex] = totalProduct / num
    }

    return outPut
}