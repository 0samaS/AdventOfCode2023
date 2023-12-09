package sols

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)

func checkAllZeroes(nums []int) bool {
	for _, num := range nums{
		if num != 0 {
			return false
		}
	}
	return true
}

func sumHistory(nums []int) int {
	if checkAllZeroes(nums) {
		return 0
	}
	var next_nums []int
	for i:=1; i < len(nums); i++{
		next_nums = append(next_nums, nums[i]-nums[i-1])
	}
	val := sumHistory(next_nums)
	return val + nums[len(nums)-1]
}

func part1Day9() int {
	res := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		nums_strs := strings.Fields(line)
		var nums []int
		for _, num_str := range nums_strs {
			num, _ := strconv.Atoi(num_str)
			nums = append(nums, num)
		}
		res += sumHistory(nums)
	}
	return res
}

func part2Day9() int {
	res := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		nums_strs := strings.Fields(line)
		var nums []int
		for i:=len(nums_strs)-1; i >= 0; i-- {
			num, _ := strconv.Atoi(nums_strs[i])
			nums = append(nums, num)
		}
		res += sumHistory(nums)
	}
	return res
}

func Day9(part string) int {
	if part == "part1" {
		return part1Day9()
	} else {
		return part2Day9()
	}
}
