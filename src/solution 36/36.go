package solution36

func isValidSudoku(board [][]byte) bool {
	// 1. Check rows

	for i := 0; i < len(board); i++ {

		row := board[i]
		all_numbers := []bool{false, false, false, false, false, false, false, false, false}
		nums := []int{}
		for _, num := range row {
			if num != '.' {
				n := int(num) - 48
				nums = append(nums, n)
				if ok := all_numbers[n-1]; ok {
					return false
				} else {
					all_numbers[n-1] = true
				}

			} else {
				nums = append(nums, 0)
			}
		}
	}

	// 2 Check columns

	for i := 0; i < len(board); i++ {
		all_numbers := []bool{false, false, false, false, false, false, false, false, false}
		nums := []int{}
		for j := 0; j < len(board); j++ {

			if board[j][i] != '.' {
				n := int(board[j][i]) - 48
				nums = append(nums, n)
				if ok := all_numbers[n-1]; ok {
					return false

				} else {
					all_numbers[n-1] = true
				}
			} else {
				nums = append(nums, 0)
			}

		}

	}

	box_numbers_vertical := []int{
		0, 3, 6,
	}
	box_numbers_horizontal := []int{
		0, 3, 6,
	}

	// 3. Check 3x3 boxes
	for _, i := range box_numbers_vertical {

		for _, j := range box_numbers_horizontal {
			all_numbers := []bool{false, false, false, false, false, false, false, false, false}
			box := board[i : i+3]
			box_horizontal := []byte{}
			for k := 0; k < 3; k++ {
				row := box[k][j : j+3]
				box_horizontal = append(box_horizontal, row...)
			}

			//check box_horizontal
			nums := []int{}
			for _, num := range box_horizontal {
				if num != '.' {

					n := int(num) - 48
					nums = append(nums, n)
					if ok := all_numbers[n-1]; ok {
						return false
					} else {
						all_numbers[n-1] = true
					}
				} else {
					nums = append(nums, 0)
				}
			}

		}
	}

	return true
}
