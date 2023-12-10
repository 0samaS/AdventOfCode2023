package sols


import (
	"bufio"
	"os"
	// "fmt"
	// "strings"
	// "strconv"
	"regexp"
)

type RowColMod struct {
	row_mod int
	col_mod int
}

type Dir int
const (
	Inv Dir = iota
	Up
	Down
	Left
	Right
)

// returns next dir
func getNextDir(field []string, r int, c int, valid_symbols map[byte]Dir) Dir {
	if r >= 0 && r < len(field) && c >= 0 && c < len(field[r]) {
		new_dir, exists := valid_symbols[field[r][c]]
		if exists {
			return new_dir
		}
	}
	return Inv
}

func findLoopLen(field []string, start_row int, start_col int) int {
	res := 1
	
	row := start_row
	col := start_col

	up_map := map[byte]Dir{'|': Up, '7': Left, 'F': Right}
	down_map := map[byte]Dir{'|': Down, 'L': Right, 'J': Left}
	left_map := map[byte]Dir{'-': Left, 'L': Up, 'F': Down}
	right_map := map[byte]Dir{'-': Right, 'J': Up, '7': Down}

	dir := Inv
	// find any valid diection to start
	if dir == Inv {
		dir = getNextDir(field, row-1, col, up_map)
		if dir != Inv {
			row--
		}
	}
	if dir == Inv {
		dir = getNextDir(field, row+1, col, down_map)
		if dir != Inv {
			row++
		}
	}
	if dir == Inv {
		dir = getNextDir(field, row, col-1, left_map)
		if dir != Inv {
			col--
		}
	}
	if dir == Inv {
		dir = getNextDir(field, row, col+1, right_map)
		if dir != Inv {
			col++
		}
	}

	for field[row][col] != field[start_row][start_col] {		
		if dir == Up {
			dir = getNextDir(field, row-1, col, up_map)
			row = row-1
		} else if dir == Down {
			dir = getNextDir(field, row+1, col, down_map)
			row = row+1
		} else if dir == Left {
			dir = getNextDir(field, row, col-1, left_map)
			col = col-1
		} else if dir == Right {
			dir = getNextDir(field, row, col+1, right_map)
			col = col+1
		}
		res++
	}
	return res
}


func part1Day10() int {	
	start_row := 0
	start_col := 0
	var field []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`S`)
		match_loc := re.FindStringIndex(line)
		if match_loc != nil {
			start_col = match_loc[0]
			start_row = len(field)
		}
		field = append(field, line)
	}
	return findLoopLen(field, start_row, start_col) / 2
}

func part2Day10() int {
	res := 0
	
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
		// line := scanner.Text()
	// }

	return res
}

func Day10(part string) int {
	if part == "part1" {
		return part1Day10()
	} else {
		return part2Day10()
	}
}
