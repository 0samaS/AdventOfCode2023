package sols

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

type RangeType struct { 
	start int
	range_len int
}

func parseInput(part_1 bool) (map[string]map[RangeType]RangeType, []string, []RangeType) {
	scanner := bufio.NewScanner(os.Stdin)

	var seeds_list []RangeType
	var mapping_key_list []string
	map_of_maps := make(map[string] map[RangeType]RangeType)
	done_with_current_mapping := true
	map_key_name := ""

	for scanner.Scan() {
		line := scanner.Text()
		if (len(line) == 0) {
			done_with_current_mapping = true
			continue
		}
		line_items := strings.Split(line, " ")
		// create seeds list
		if line_items[0] == "seeds:" {
			// seed range for part 2
			if part_1 {
				for i := 1; i < len(line_items); i++ {
					seed_num, _ := strconv.Atoi(line_items[i])
					seeds_list = append(seeds_list, RangeType{seed_num, 0})
				}
			} else {
				for i := 1; i < len(line_items); i+=2 {
					seed_num_start, _ := strconv.Atoi(line_items[i])
					seed_num_range, _ := strconv.Atoi(line_items[i+1])
					seeds_list = append(seeds_list, RangeType{seed_num_start, seed_num_range})
				}
			}
			continue
		}
		if done_with_current_mapping {
			split_mapping_names := strings.Split(line_items[0], "-")
			// backwards mapping for part 2
			if part_1 {
				map_key_name = split_mapping_names[0]
			} else {
				map_key_name = split_mapping_names[2]
			}
			mapping_key_list = append(mapping_key_list, map_key_name)
			map_of_maps[map_key_name] = make(map[RangeType]RangeType)

			done_with_current_mapping = false
			continue
		}
		dest_start, _ := strconv.Atoi(line_items[0])
		src_start, _ := strconv.Atoi(line_items[1])
		range_len, _ := strconv.Atoi(line_items[2])

		// flip src and dest for part 2
		if part_1 {
			map_of_maps[map_key_name][RangeType{src_start, range_len}] = RangeType{dest_start, range_len}
		} else {
			map_of_maps[map_key_name][RangeType{dest_start, range_len}] = RangeType{src_start, range_len}
		}
	}
	return map_of_maps, mapping_key_list, seeds_list
}

func translatableToSeed(map_of_maps *map[string]map[RangeType]RangeType, mapping_key_list *[]string, seeds_list *[]RangeType, tester int) bool {
	intermediate := tester
	for i:=len(*mapping_key_list)-1; i >= 0; i-- {
		key := (*mapping_key_list)[i]
		for src_info, dest_info:= range (*map_of_maps)[key] {
			// this can be translated
			if intermediate >= src_info.start && intermediate < src_info.start + src_info.range_len {
				intermediate = (intermediate - src_info.start + dest_info.start)
				break
			}
		}
	}
	for _, seed_info := range *seeds_list {
		if intermediate > seed_info.start && intermediate < seed_info.start + seed_info.range_len {
			return true
		}
	}
	return false
}

func part1Day5() int {
	map_of_maps, mapping_key_list, seeds_list := parseInput(true)
	min_location := math.MaxInt32
	for _, seed := range seeds_list {
		intermediate := seed.start
		for _, key := range mapping_key_list {
			for src_info, dest_info:= range map_of_maps[key] {
				// this can be translated
				if intermediate >= src_info.start && intermediate < src_info.start + src_info.range_len {
					intermediate = (intermediate - src_info.start + dest_info.start)
					break
				}
			}
		}
		if intermediate < min_location {
			min_location = intermediate
		}
	}
	return min_location
}

func part2Day5() int {
	map_of_maps, mapping_key_list, seeds_list := parseInput(false)
	location := 0
	for !translatableToSeed(&map_of_maps, &mapping_key_list, &seeds_list, location) {
		location++;
	}
	return location
}

func Day5(part string) int {
	if part == "1" {
		return part1Day5()
	} else {
		return part2Day5()
	}
}
