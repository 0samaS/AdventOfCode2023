package sols

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
	"math"
)

type RangeType struct { 
	start int
	range_len int
} 

func part1Day5() int {
	scanner := bufio.NewScanner(os.Stdin)

	var seeds_list []int
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
			for i := 1; i < len(line_items); i++ {
				seed_num, _ := strconv.Atoi(line_items[i])
				seeds_list = append(seeds_list, seed_num)
			}
			continue
		}
		if done_with_current_mapping {
			split_mapping_names := strings.Split(line_items[0], "-")
			map_key_name = split_mapping_names[0]
			mapping_key_list = append(mapping_key_list, map_key_name)
			map_of_maps[map_key_name] = make(map[RangeType]RangeType)
			done_with_current_mapping = false
			continue
		}
		dest_start, _ := strconv.Atoi(line_items[0])
		src_start, _ := strconv.Atoi(line_items[1])
		range_len, _ := strconv.Atoi(line_items[2])

		map_of_maps[map_key_name][RangeType{src_start, range_len}] = RangeType{dest_start, range_len}
	}
	min_location := math.MaxInt32
	for _, seed := range seeds_list {
		intermediate := seed
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
			for i := 1; i < len(line_items); i+=2 {
				seed_num_start, _ := strconv.Atoi(line_items[i])
				seed_num_range, _ := strconv.Atoi(line_items[i+1])
				seeds_list = append(seeds_list, RangeType{seed_num_start, seed_num_range})
			}
			continue
		}
		if done_with_current_mapping {
			// fmt.Printf("%s\n", line)
			if line_items[1] != "map:" {
				fmt.Printf("Something has gone terribly wrong!")
			}
			split_mapping_names := strings.Split(line_items[0], "-")

			map_key_name = split_mapping_names[0]
			mapping_key_list = append(mapping_key_list, map_key_name)
			map_of_maps[map_key_name] = make(map[RangeType]RangeType)
			// fmt.Printf("\n\n%s key name %d \n", map_key_name, len(map_of_maps))
			done_with_current_mapping = false
			continue
		}
		dest_start, _ := strconv.Atoi(line_items[0])
		src_start, _ := strconv.Atoi(line_items[1])
		range_len, _ := strconv.Atoi(line_items[2])

		map_of_maps[map_key_name][RangeType{src_start, range_len}] = RangeType{dest_start, range_len}
	}
	min_location := math.MaxInt32

	for _, seed_info := range seeds_list {
		// fmt.Printf("%d, %d\n", seed_info.start, seed_info.range_len)
		for seed:=seed_info.start; seed < seed_info.start + seed_info.range_len; seed++ {
			intermediate := seed
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

	}	

	return min_location
}

func Day5(part string) int {
	if part == "1" {
		return part1Day5()
	} else {
		return part2Day5()
	}
}
