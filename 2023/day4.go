package main

import (
	"math"
	"strconv"
	"strings"
)

func day4part1() int {
	input := strings.Split(readInput("in/day4.in"), "\n")
	sum := 0

	for _, line := range input {
		points := countPoints(line)
		sum += int(math.Pow(2, float64(points-1)))
	}

	return sum
}

func isWinning(winnning []int, number int) bool {
	for _, v := range winnning {
		if number == v {
			return true
		}
	}
	return false
}

func countPoints(line string) (points int) {
	var winning []int
	numbers := strings.Split(strings.Split(line, ":")[1], "|")
	for _, v := range strings.Split(numbers[0], " ") {
		if v != "" {
			number, _ := strconv.Atoi(v)
			winning = append(winning, number)
		}
	}
	for _, v := range strings.Split(numbers[1], " ") {
		if v != "" {
			number, _ := strconv.Atoi(v)
			if isWinning(winning, number) {
				points++
			}
		}
	}
	return
}

func day4part2() int {
	input := strings.Split(readInput("in/day4.in"), "\n")
	sum := 0
	points := map[int]int{}
	for i, line := range input {
		points[i] = countPoints(line)
	}

	for i := range input {
		sum += scratch(points, i)
	}

	return sum
}

func scratch(points map[int]int, start int) int {
	sum := 1
	for i := 0; i < points[start]; i++ {
		sum += scratch(points, start+i+1)
	}
	return sum
}
