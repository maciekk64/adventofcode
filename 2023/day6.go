package main

import (
	"strconv"
	"strings"
)

func day6() int {
	races := parseInput6part2()
	ways := 1

	for _, race := range races {
		ways *= waysToWin(race)
	}

	return ways
}

func waysToWin(race [2]int) (sum int) {
	for i := 1; i < race[0]; i++ {
		distance := i * (race[0] - i)
		if distance > race[1] {
			sum++
		}
	}
	return
}

func parseInput6part1() (races [][2]int) {
	input := strings.Split(readInput("in/day6.in"), "\n")
	x := strings.Fields(strings.Split(input[0], ":")[1])
	y := strings.Fields(strings.Split(input[1], ":")[1])

	for i := range x {
		time, _ := strconv.Atoi(x[i])
		distance, _ := strconv.Atoi(y[i])
		races = append(races, [2]int{time, distance})
	}

	return
}

func parseInput6part2() (races [][2]int) {
	input := strings.Split(readInput("in/day6.in"), "\n")
	x := strings.ReplaceAll(strings.Split(input[0], ":")[1], " ", "")
	y := strings.ReplaceAll(strings.Split(input[1], ":")[1], " ", "")

	time, _ := strconv.Atoi(x)
	distance, _ := strconv.Atoi(y)
	races = append(races, [2]int{time, distance})

	return
}
