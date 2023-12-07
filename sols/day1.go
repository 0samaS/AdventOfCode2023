package sols

import (
	"bufio"
	"os"
	"unicode"
	"regexp"
	// "fmt"
)

func part1Day1() int {
	res := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var first, current int = 0, 0
		var found_first = false
		for _, c := range line {
			if unicode.IsDigit(c){
				current = int(c - '0')
				if !found_first{
					first = current
					found_first = true
				}
			}
		}
		res += (first*10 + current)
	}

	return res
}

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return
}

func part2Day1() int {
	digitsMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	forward_re := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	back_re := regexp.MustCompile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
	res := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		forward_matches := forward_re.FindAllString(line, -1)
		back_matches := back_re.FindAllString(Reverse(line), -1)
		res += (digitsMap[forward_matches[0]]*10 + digitsMap[Reverse(back_matches[0])])
	}
	return res
}

func Day1(part string) int {
	if part == "part1" {
		return part1Day1()
	} else {
		return part2Day1()
	}
}
