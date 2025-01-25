
func isValidSudoku(board [][]byte) bool {
	x := 0
	for x < 9 {
		isValid := isValideIndexRow(x, board)
		if !isValid {
			return false
		}

		isValid = validateIndexColumns(x, board)
		if !isValid {
			return false
		}
		x++
	}

	return validateQuarters(board)
}

func isValideIndexRow(index int, board [][]byte) bool {
	row := board[index]
	foundNums := make(map[byte]bool, 9)

	for _, num := range row {
		if num == 46 {
			continue
		}

		if !foundNums[num] {
			foundNums[num] = true
		} else {
			return false
		}
	}

	return true
}

func validateIndexColumns(index int, board [][]byte) bool {
	foundNums := make(map[byte]bool, 9)

	for _, row := range board {
		if row[index] == 46 {
			continue
		}

		if !foundNums[row[index]] {
			foundNums[row[index]] = true
		} else {
			return false
		}

	}

	return true
}

func validateQuarters(board [][]byte) bool {
	sets := map[int][]byte{}
	rowSet := 1
	for rowIndex, row := range board {
		columnSet := 1
		for columnIndex, column := range row {
			setName := ((rowSet) * 10) + (columnSet)
			sets[setName] = append(sets[setName], column)
			if (columnIndex+1)%3 == 0 {
				columnSet++
			}

		}

		if (rowIndex+1)%3 == 0 {
			rowSet++
		}
	}

	fmt.Println(sets)

	foundNums := make(map[byte]bool, 9)

	for _, set := range sets {

		for _, num := range set {
			if num == 46 {
				continue
			}

			if !foundNums[num] {
				foundNums[num] = true
			} else {
				return false
			}
		}
		foundNums = make(map[byte]bool, 9)

	}
	return true

}
