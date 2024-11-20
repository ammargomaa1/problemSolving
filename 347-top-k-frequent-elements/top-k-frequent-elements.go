func topKFrequent(nums []int, k int) []int {
	mapOfElements := map[int]int{}
	sliceOfCount := make([][]int, len(nums)+1)
	for _, num := range nums {
		mapOfElements[num]++
	}
	
	for num, appearances := range mapOfElements {
		
		sliceOfCount[appearances] = append(sliceOfCount[appearances], num)
		
	}
	sliceOfAnswer := []int{}

	for i := len(sliceOfCount)-1;  i > 0; i-- {
		sliceOfAnswer = append(sliceOfAnswer, sliceOfCount[i]...)
		if len(sliceOfAnswer) >= k {
			return sliceOfAnswer[:k]
		}
	}

	
	return sliceOfAnswer
}