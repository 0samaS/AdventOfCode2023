package sols

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

func parseLoop(line string, process func(string, int) bool) {
	line = line[8: len(line)]
	line = strings.ReplaceAll(line, ",", "")
	rounds := strings.Split(line, ";")
	for _, s := range rounds{
		s = strings.TrimSpace(s)
		round_info := strings.Split(s, " ")
		for i := 0; i < len (round_info); i+=2 {
			if num, err := strconv.Atoi(round_info[i]); err == nil {
				color := round_info[i+1]
				process(color, num)
			}
		}
	}
}

func part1Day2() int {
	res := 0
	constraints := map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		possible := true

		line := scanner.Text()
		split_game := strings.Split(line, ":")
		game_num, _ := strconv.Atoi(split_game[0][5:len(split_game[0])]);

		parseLoop(line, func(color string, num int) bool {
										if constraints[color] < num {
											possible = false
											return false
										}
										return true
		})
		if possible {
			res += game_num
		}
	}
	return res
}

func part2Day2() int {
	res := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		max_cubes := map[string]int {
			"red": 0,
			"green": 0,
			"blue": 0,
		}

		line := scanner.Text()
		parseLoop(line, func(color string, num int) bool {
										if num > max_cubes[color]{
											max_cubes[color] = num
										}
										return true
		})
		res += (max_cubes["red"] * max_cubes["green"] * max_cubes["blue"])
	}
	return res
}

func Day2(part string) int {
	if part == "1" {
		return part1Day2()
	} else {
		return part2Day2()
	}
}
