func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // Step 1: Add all numbers to a map for O(1) lookups
    existingNumbers := make(map[int]bool, len(nums))
    for _, num := range nums {
        existingNumbers[num] = true
    }

    maxSeq := 0

    // Step 2: Process each number
    for number := range existingNumbers {
        // Check if this number is the start of a sequence
        if !existingNumbers[number-1] {
            currentNum := number
            seqLen := 1

            // Count the length of the sequence
            for existingNumbers[currentNum+1] {
                currentNum++
                seqLen++
            }

            // Update the maximum sequence length
            if seqLen > maxSeq {
                maxSeq = seqLen
            }
        }
    }

    return maxSeq
}
