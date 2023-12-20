package main

import (
	"regexp"
	"strings"
)

func day8part1() (i int) {
	input, nav := parseInput8()

	node := "AAA"
	for l := len(nav); node != "ZZZ"; i++ {
		node = next(input[node], nav[i%l])
	}

	return
}

func next(node struct{ left, right string }, num byte) string {
	if num == 76 {
		return node.left
	} else {
		return node.right
	}
}

func parseInput8() (map[string]struct{ left, right string }, string) {
	input := map[string]struct{ left, right string }{}
	lines := strings.Split(readInput("in/day8.in"), "\n")
	nav := lines[0]

	for _, v := range lines[2:] {
		r := regexp.MustCompile(`[A-Z]{3}`)
		matches := r.FindAllStringSubmatch(v, -1)
		input[strings.Join(matches[0], "")] = struct{ left, right string }{
			strings.Join(matches[1], ""), strings.Join(matches[2], ""),
		}
	}

	return input, nav
}

// func day8part2() (i int) {
// 	input, nav := parseInput8()
// 	nodes := []string{}
// 	for i := range input {
// 		if i[2] == 65 {
// 			nodes = append(nodes, i)
// 		}
// 	}
//
// 	for l := len(nav); !allEndsWithZ(nodes); i++ {
// 		for j := 0; j < 6; j++ {
// 			nodes[j] = next(input[nodes[j]], nav[i%l])
// 		}
// 	}
//
// 	return
// }
//
// func allEndsWithZ(nodes []string) bool {
// 	for _, node := range nodes {
// 		if node[2] != 90 {
// 			return false
// 		}
// 	}
// 	return true
// }
