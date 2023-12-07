package sols

import (
	"bufio"
	"os"
	"strings"
	// "fmt"
	"strconv"
	"math"
)

func getNumInCommon(winning_nums []string, our_nums []string) int {
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
	return num_in_common
}

func part1Day4() int {
	res := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		split_card := strings.Split(line, ":")
		split_card = strings.Split(split_card[1], "|")

		winning_nums := strings.Fields(strings.TrimSpace(split_card[0]))
		our_nums := strings.Fields(strings.TrimSpace(split_card[1]))
		num_in_common := getNumInCommon(winning_nums, our_nums)

		res += int(math.Pow(2, float64(num_in_common - 1)))
	}
	return res
}

func part2Day4() int {
	res := 0
	
	// stores how many copies of this card num we have
	copies_map := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		split_card := strings.Split(line, ":")
		card_num, _ := strconv.Atoi(strings.TrimSpace(split_card[0][4:len(split_card[0])]));
		split_card = strings.Split(split_card[1], "|")

		winning_nums := strings.Fields(strings.TrimSpace(split_card[0]))
		our_nums := strings.Fields(strings.TrimSpace(split_card[1]))
		num_in_common := getNumInCommon(winning_nums, our_nums)

		// include original in copies
		copies_map[card_num] += 1
		for i := 0; i < copies_map[card_num]; i++ {
			for j := card_num+1; j < card_num+1+num_in_common; j++ {
				copies_map[j] += 1
			}
		}
	}
	for _, num_copies := range copies_map {
		res += num_copies
	}
	return res
}

func Day4(part string) int {
	if part == "part1" {
		return part1Day4()
	} else {
		return part2Day4()
	}
}
