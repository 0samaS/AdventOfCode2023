package sols


import (
	"bufio"
	"os"
	// "fmt"
	// "strings"
	// "strconv"
	"regexp"
)

// var dist_map = map[]

// func dist(v1 int, v2 int, u1 int, u2 int) {

// }

// func APSP(expanded_universe []string) {
// 	var dist []
// 	for v1:=0; v1<len(expanded_universe); v1++ {
// 		for u1:=0; u1<len(expanded_universe); u1++ {
// 			for v2:=0; v2<len(expanded_universe); v2++ {
// 				for u2:=0; u2<len(expanded_universe); u2++ {
// 					if v1 == u1 && v2 == u2 {
// 						dist[Node{v1,}]
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Galaxy struct {
	row int
	col int
}

type GalaxyPair struct {
	first int
	second int
}

func findManhattanDist(x Galaxy, y Galaxy) int {
	return abs(x.row - y.row) + abs(x.col - y.col)
}

func part1Day11() int {
	res := 0
	var universe []string
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := scanner.Text()
		universe = append(universe, line)
	}
	col_info := make([]int, len(universe[0]))
	for i:=0; i<len(universe); i++ {
		for j:=0; j<len(universe[i]); j++ {
			if universe[i][j] == '#' {
				col_info[j] += 1
			} else {
				col_info[j] += 0
			}
		}
	}

	var expanded_universe []string
	re := regexp.MustCompile(`#`)

	for _, row := range universe {
		new_row := row
		buffer_col := 0

		for col_num, num_galaxies := range col_info {

			if num_galaxies == 0 {
				new_row = new_row[0:col_num+buffer_col] + string('.') + new_row[col_num+buffer_col:len(new_row)]
				buffer_col++
			}
		}
		expanded_universe = append(expanded_universe, new_row)
		matches := re.FindString(new_row)
		// add extra row if empty
		if len(matches) == 0 {
			expanded_universe = append(expanded_universe, new_row)
		}
	}

	var galaxies []Galaxy

	for row_num, line := range expanded_universe {
		for col_num, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Galaxy{row_num, col_num})
			}
		}
	}

	dist_map := make(map[GalaxyPair]int)
	for i:=0; i<len(galaxies); i++ {
		for j:=i+1; j<len(galaxies); j++ {
			galaxy_pair := GalaxyPair{i, j}
			dist_map[galaxy_pair] = findManhattanDist(galaxies[i], galaxies[j])
			res += dist_map[GalaxyPair{i, j}]
		}
	}

	return res
}

func part2Day11() int {
	res := 0
	
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
		// line := scanner.Text()
	// }

	return res
}

func Day11(part string) int {
	if part == "part1" {
		return part1Day11()
	} else {
		return part2Day11()
	}
}
