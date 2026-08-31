package ai

func ChooseMove(cells []string) (int, string) {
	if len(cells) != 9 {
		return 0, ""
	}

	move := findWinningMove(cells, "O")
	if move != 0 {
		return move, "win"
	}

	move = findWinningMove(cells, "X")
	if move != 0 {
		return move, "block"
	}

	if cells[4] == "" {
		return 5, "center"
	}

	corners := []int{1, 3, 7, 9}
	for _, cell := range corners {
		if cells[cell-1] == "" {
			return cell, "corner"
		}
	}

	sides := []int{2, 4, 6, 8}
	for _, cell := range sides {
		if cells[cell-1] == "" {
			return cell, "side"
		}
	}

	return 0, ""
}

func findWinningMove(cells []string, mark string) int {
	for cell := 1; cell <= 9; cell++ {
		index := cell - 1

		if cells[index] != "" {
			continue
		}

		cells[index] = mark

		if hasWin(cells, mark) {
			cells[index] = ""
			return cell
		}

		cells[index] = ""
	}

	return 0
}

func hasWin(cells []string, mark string) bool {
	lines := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, line := range lines {
		first := line[0]
		second := line[1]
		third := line[2]

		if cells[first] == mark &&
			cells[second] == mark &&
			cells[third] == mark {
			return true
		}
	}

	return false
}
