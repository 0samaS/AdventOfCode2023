package sols


import (
	"bufio"
	"os"
)

type Galaxy struct {
	row int
	col int
}

func findExpandedManhattanDist(x Galaxy, y Galaxy, row_info[]int, col_info[]int, len_expansion int) int {
	min_row := 0
	max_row := 0
	min_col := 0
	max_col := 0
	if x.row < y.row {
		min_row = x.row
		max_row = y.row

	} else if y.row < x.row {
		min_row = y.row
		max_row = x.row
	}
	if x.col < y.col {
		min_col = x.col
		max_col = y.col

	} else if y.col < x.col {
		min_col = y.col
		max_col = x.col
	}

	row_dist := 0
	for min_row < max_row {
		if row_info[min_row] == 0 {
			row_dist += len_expansion
		} else {
			row_dist += 1
		}
		min_row++
	}

	col_dist := 0
	for min_col < max_col {
		if col_info[min_col] == 0 {
			col_dist += len_expansion
		} else {
			col_dist += 1
		}
		min_col++
	}
	return row_dist + col_dist
}

func parseAndCalculate(len_expansion int) int{
	res := 0
	var universe []string
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := scanner.Text()
		universe = append(universe, line)
	}
	col_info := make([]int, len(universe[0]))
	row_info := make([]int, len(universe))
	for i:=0; i<len(universe); i++ {
		for j:=0; j<len(universe[i]); j++ {
			if universe[i][j] == '#' {
				col_info[j] += 1
				row_info[i] += 1
			} else {
				col_info[j] += 0
				row_info[i] += 0
			}
		}
	}
	var galaxies []Galaxy
	for row_num, line := range universe {
		for col_num, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Galaxy{row_num, col_num})
			}
		}
	}
	for i:=0; i<len(galaxies); i++ {
		for j:=i+1; j<len(galaxies); j++ {
			res += findExpandedManhattanDist(galaxies[i], galaxies[j], row_info, col_info, len_expansion)
		}
	}
	return res
}

func part1Day11() int {
	return parseAndCalculate(2)
}

func part2Day11() int {
	return parseAndCalculate(1000000)
}

func Day11(part string) int {
	if part == "part1" {
		return part1Day11()
	} else {
		return part2Day11()
	}
}
