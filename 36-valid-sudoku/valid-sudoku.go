func isValidSudoku(board [][]byte) bool {
	columns := make([]map[byte]bool,9)

	rows :=  make([]map[byte]bool,9)

	quarters := make(map[string]map[byte]bool, 9)

	for x, row := range board{
		quarterXPosition := x/3
		for y, cell := range row{
			if string(cell) == "." {
				continue
			}

			if columns[y] == nil{
				columns[y] = make(map[byte]bool, 9)
			}

			if columns[y][cell]  {
				return false
			}
			columns[y][cell] = true

			if rows[x] == nil {
				rows[x] = make(map[byte]bool, 9)
			}

			if rows[x][cell] {
				return false
			}
			rows[x][cell] = true

			quarterYPosition := y/3
			if quarters[string(quarterXPosition) + string(quarterYPosition)] == nil {
				quarters[string(quarterXPosition) + string(quarterYPosition)] = make(map[byte]bool, 9)
			}
			
			if quarters[string(quarterXPosition) + string(quarterYPosition)][cell] {
				return false
			}
			quarters[string(quarterXPosition) + string(quarterYPosition)][cell] = true
		}
	}

	return true
}