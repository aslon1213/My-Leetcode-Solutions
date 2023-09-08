package main

func main() {

	generate(5)

}

func generate(numRows int) [][]int {

	if numRows == 1 {
		return [][]int{{1}}
	}

	output := [][]int{
		{1}, {1, 1},
	}

	for i := 2; i < numRows; i++ {
		row := []int{}
		length_row := i + 1
		for j := 0; j < length_row; j++ {
			if j == 0 || j == length_row-1 {
				row = append(row, 1)
			} else {
				row = append(row, output[i-1][j-1]+output[i-1][j])
			}
		}
		output = append(output, row)
	}
	return output

}
