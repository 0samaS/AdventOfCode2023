package sols

import (
	"bufio"
	"os"
	"strings"
	// "fmt"
	"strconv"
	"math"
)

func part1Day4() int {
	res := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		split_card := strings.Split(line, ":")
		split_card = strings.Split(split_card[1], "|")

		winning_nums := strings.Fields(strings.TrimSpace(split_card[0]))
		our_nums := strings.Fields(strings.TrimSpace(split_card[1]))

		our_nums_map := make(map[int]int)
		winning_nums_map := make(map[int]int)

		for _, num_str := range our_nums {
			num, _ := strconv.Atoi(num_str)
			our_nums_map[num] += 1
		}
		for _, num_str := range winning_nums {
			num, _ := strconv.Atoi(num_str)

			winning_nums_map[num] += 1
		}

		// check if winning numbers appear in our set
		num_in_common := 0
		for num, _ := range our_nums_map {

			_, exists := winning_nums_map[num]
			if exists {
				num_in_common += 1
			}
		}
		res += int(math.Pow(2, float64(num_in_common - 1)))
	}
	return res
}

func part2Day4() int {
	res := 0

	return res
}

func Day4(part string) int {
	if part == "1" {
		return part1Day4()
	} else {
		return part2Day4()
	}
}
