package sols

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"sort"
)

type HandType int
const (
	HighCard	HandType = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

var card_ranking = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

func getHandType(hand string) HandType {
	card_count_map := make(map[rune]int)
	for _, c := range hand {
		card_count_map[rune(c)] += 1
	}
	one_pair   := 0
	three_kind := 0 
	four_kind  := 0
	five_kind  := 0

	for _, count := range card_count_map {
		if count == 2 {
			one_pair++
		} else if count == 3 {
			three_kind++;
		} else if count == 4 {
			four_kind++;
		} else if count == 5 {
			five_kind++;
		}
	}
	if five_kind == 1 {
		return FiveKind
	} else if four_kind == 1 {
		return FourKind
	} else if three_kind == 1 && one_pair == 1 {
		return FullHouse
	} else if three_kind == 1 {
		return ThreeKind
	} else if one_pair == 2 {
		return TwoPair
	} else if one_pair == 1 {
		return OnePair
	}

	return HighCard
}

// implement sort interface
type HandComp []string

func (a HandComp) Len() int { 
	return len(a) 
}
func (a HandComp) Less(i, j int) bool {
	hand_type_i := getHandType(a[i])
	hand_type_j := getHandType(a[j])
	if hand_type_i == hand_type_j {
		for idx:=0; idx < len(a); idx++ {
			card_rank_i:= card_ranking[rune(a[i][idx])]
			card_rank_j:= card_ranking[rune(a[j][idx])]
			if card_rank_i < card_rank_j {
				return true
			} else if card_rank_i > card_rank_j {
				return false
			} else {
				continue
			}
		}
	} 
	return hand_type_i < hand_type_j
}
func (a HandComp) Swap(i, j int) {
	 a[i], a[j] = a[j], a[i] 
}

func part1Day7() int {
	res := 0

	hand_bid_map := make(map[string]int)
	var hand_vector []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line_items := strings.Fields(scanner.Text())
		hand_bid_map[line_items[0]], _ = strconv.Atoi(line_items[1])
		hand_vector = append(hand_vector, line_items[0])
		// fmt.Printf("%s, %s\n", line_items[0], line_items[1])
	}

	sort.Sort(HandComp(hand_vector))

	for rank, hand := range hand_vector {
		fmt.Printf("%s\n", hand)
		res += ((rank+1) * hand_bid_map[hand])
	}

	return res
}

func part2Day7() int {
	res := 0

	return res
}

func Day7(part string) int {
	if part == "1" {
		return part1Day7()
	} else {
		return part2Day7()
	}
}
