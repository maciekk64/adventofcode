package main

import (
	"regexp"
	"strconv"
	"strings"
)

func day2part1() int {
	input := strings.Split(readInput("in/day2.in"), "\n")
	r := regexp.MustCompile(`([0-9]+) (red|green|blue)`)
	maxValues := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sum := 0

	for i, v := range input {
		matches := r.FindAllStringSubmatch(v, -1)
		if isPossible(matches, maxValues) {
			sum += i + 1
		}
	}

	return sum
}

func isPossible(matches [][]string, maxValues map[string]int) bool {
	for _, v := range matches {
		number, _ := strconv.Atoi(v[1])
		color := v[2]
		if number > maxValues[color] {
			return false
		}
	}
	return true
}

func day2part2() int {
	input := strings.Split(readInput("in/day2.in"), "\n")
	r := regexp.MustCompile(`([0-9]+) (red|green|blue)`)
	sum := 0

	for _, v := range input {
		matches := r.FindAllStringSubmatch(v, -1)
		sum += minCubes(matches)
	}

	return sum
}

func minCubes(matches [][]string) int {
	numbers := make(map[string]int)
	for _, v := range matches {
		number, _ := strconv.Atoi(v[1])
		color := v[2]
		if number > numbers[color] {
			numbers[color] = number
		}
	}
	return numbers["red"] * numbers["green"] * numbers["blue"]
}
