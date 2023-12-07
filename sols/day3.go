package sols

import (
	"bufio"
	"os"
	"unicode"
	"strconv"
)

type MapLocation struct { 
	row    int 
	col    int 
} 

type PartInfo struct {
	row int
	starting_idx int
	len int
}

func checkStar(all_lines []string, row int, col int) bool {
	if (row >= 0 && row < len(all_lines) && col >= 0 && col < len(all_lines[row])) {
		if all_lines[row][col] == '*' {
			return true
		}
	}
	return false
}

func checkForStars(gear_map map[MapLocation]map[PartInfo]int, 
												all_lines[]string, row int, starting_idx int, len int, current_num int) {

	for i:= starting_idx; i < starting_idx + len; i++ {
		for r := row-1; r <= row+1; r++ {
			for c:= i-1; c <= i+1; c++ {
				if checkStar(all_lines, r, c) {
					val, ok := gear_map[MapLocation{r, c}]
					if !ok {
						val = make(map[PartInfo]int)
					}
					val[PartInfo{row, starting_idx, len}] = current_num
					gear_map[MapLocation{r, c}] = val
				}
			}
		}
	}
}

func checkIndex(all_lines []string, row int, col int) bool {
	if (row >= 0 && row < len(all_lines) && col >= 0 && col < len(all_lines[row])) {
		if all_lines[row][col] != '.' && !unicode.IsDigit(rune(all_lines[row][col])) {
			return true
		}
	}
	return false
}

func checkBoundary(all_lines[]string, row int, starting_idx int, len int) bool{
	is_part := false
	for i:= starting_idx; i < starting_idx + len; i++ {
		for r := row-1; r <= row+1; r++ {
			for c:= i-1; c <= i+1; c++ {
				is_part = is_part || checkIndex(all_lines, r, c)
			}
		}
	}
	return is_part
}

func part1Day3() int {
	res := 0
	var all_lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		all_lines = append(all_lines, line)
	}

	for row, line := range all_lines {
		starting_idx := -1
		var current_num_str string
		for i := 0; i < len(line); i++ {
			c := rune(line[i])

			if unicode.IsDigit(c) {
				if starting_idx == -1 {
					starting_idx = i
				}
				current_num_str += string(c)

			} else {
				// we have a number to prcoess
				if starting_idx != -1 {
					if checkBoundary(all_lines, row, starting_idx, len(current_num_str)) {
						part_num, _ := strconv.Atoi(current_num_str);
						res += part_num
					}
					// reset vars for next iteration
					starting_idx = -1
					current_num_str = ""
				}
			}
		}
		// we have a number to prcoess after the loop
		if starting_idx != -1 {
			if checkBoundary(all_lines, row, starting_idx, len(current_num_str)) {
				part_num, _ := strconv.Atoi(current_num_str);
				res += part_num
			}
			// reset vars for next iteration
			starting_idx = -1
			current_num_str = ""
		}
	}
	return res
}

func part2Day3() int {
	res := 0
	var all_lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		all_lines = append(all_lines, line)
	}

	gear_map := make(map[MapLocation]map[PartInfo]int)

	for row, line := range all_lines {
		starting_idx := -1
		var current_num_str string
		for i := 0; i < len(line); i++ {
			c := rune(line[i])

			if unicode.IsDigit(c) {
				if starting_idx == -1 {
					starting_idx = i
				}
				current_num_str += string(c)

			} else {
				// we have a number to prcoess
				if starting_idx != -1 {

					current_num, _ := strconv.Atoi(current_num_str)
					checkForStars(gear_map, all_lines, row, starting_idx, len(current_num_str), current_num) 

					// reset vars for next iteration
					starting_idx = -1
					current_num_str = ""
				}
			}
		}
		// we have a number to prcoess after the loop
		if starting_idx != -1 {

			current_num, _ := strconv.Atoi(current_num_str)
			checkForStars(gear_map, all_lines, row, starting_idx, len(current_num_str), current_num) 

			// reset vars for next iteration
			starting_idx = -1
			current_num_str = ""
		}
	}

	// process all the gears and get their ratios
	for _, map_adjacent_parts := range gear_map {
		if len(map_adjacent_parts) == 2 {

			intermediate := 1
			for _, num := range map_adjacent_parts {
				intermediate *= num
			}
			res += intermediate
		}
	}
	return res
}

func Day3(part string) int {
	if part == "part1" {
		return part1Day3()
	} else {
		return part2Day3()
	}
}
