package sols


import (
	"bufio"
	"os"
	"strings"
)

type Node struct {
	left string
	right string
}

func gcd(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func lcm(a, b int) int {
	return a * (b / (gcd(a, b)))
}

func part1Day8() int {
	node_map := make(map[string]Node)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	instructions := scanner.Text()
	// consume empty line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, " = ")
		child_nodes := strings.Split(nodes[1], ", ")
		left_child := child_nodes[0][1:len(child_nodes[0])]
		right_child := child_nodes[1][0:len(child_nodes[1])-1]

		parent := nodes[0]
		node_map[parent] = Node{left_child, right_child}
	}
	num_steps := 0
	instruction_idx := 0
	current := "AAA"

	for current != "ZZZ" {
		// left
		if rune(instructions[instruction_idx]) == rune('L') {
			current = node_map[current].left
		} else { // right
			current = node_map[current].right
		}
		instruction_idx = (instruction_idx + 1) % len(instructions)
		num_steps++
	}

	return num_steps
}

func part2Day8() int {
	node_map := make(map[string]Node)
	var current []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	instructions := scanner.Text()
	// consume empty line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, " = ")
		child_nodes := strings.Split(nodes[1], ", ")
		left_child := child_nodes[0][1:len(child_nodes[0])]
		right_child := child_nodes[1][0:len(child_nodes[1])-1]

		parent := nodes[0]
		node_map[parent] = Node{left_child, right_child}
		if rune(parent[2]) == rune('A') {
			current = append(current, parent)
		}
	}

	res := 1
	for i:=0; i < len(current); i++ {
		instruction_idx := 0
		path_len := 0
		for current[i][2] != 'Z' {
			// left
			if rune(instructions[instruction_idx]) == rune('L') {
				current[i] = node_map[current[i]].left
			} else { // right
				current[i] = node_map[current[i]].right
			}
			instruction_idx = (instruction_idx + 1) % len(instructions)
			path_len++
		}
		res = lcm(path_len, res)
	}
	return res
}

func Day8(part string) int {
	if part == "part1" {
		return part1Day8()
	} else {
		return part2Day8()
	}
}
