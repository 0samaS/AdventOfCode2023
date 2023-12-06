package sols

import (
	"bufio"
	"os"
	// "fmt"
	"strings"
	"strconv"
)

type RaceInfo struct {
	time int
	distance int
}

func numWays(race RaceInfo) int {
	num_ways := 0
	for hold_time:=0; hold_time < race.time; hold_time++ {
		if (hold_time * (race.time - hold_time)) > race.distance {
			num_ways++
		}
	}
	return num_ways
}

func part1Day6() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 
	time_items := strings.Fields(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]));

	scanner.Scan() 
	distance_items := strings.Fields(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]));

	res := 1
	for i:=0; i < len(time_items); i++ {
		time_num, _ := strconv.Atoi(time_items[i])
		distance_num, _ := strconv.Atoi(distance_items[i])
		num_ways := numWays(RaceInfo{time_num, distance_num})
		res *= num_ways
	}

	return res
}

func part2Day6() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 
	time_items := strings.Fields(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]));
	time_str := ""
	for _, str := range time_items {
		time_str += str
	}
	time_num, _ := strconv.Atoi(time_str)

	scanner.Scan() 
	distance_items := strings.Fields(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]));

	distance_str := ""
	for _, str := range distance_items {
		distance_str += str
	}
	distance_num, _ := strconv.Atoi(distance_str)

	res := numWays(RaceInfo{time_num, distance_num})

	return res
}

func Day6(part string) int {
	if part == "1" {
		return part1Day6()
	} else {
		return part2Day6()
	}
}
