package main

import (
	"math"
	"strconv"
	"strings"
)

func day5part1() int {
	seeds, numbers := parseInput5()
	location := math.MaxInt

	for _, seed := range seeds {
		thing := seed
		for _, v := range numbers {
			thing = convertThingToTheNextThing(thing, v)
		}
		if thing < location {
			location = thing
		}
	}
	return location
}

func parseInput5() ([]int, [7][][]int) {
	input := strings.Split(readInput("in/day5.in"), "\n\n")

	var seeds []int
	for _, v := range strings.Split(strings.TrimLeft(strings.Split(input[0], ":")[1], " "), " ") {
		seed, _ := strconv.Atoi(v)
		seeds = append(seeds, seed)
	}

	var numbers [7][][]int
	for i, convert := range input[1:] {
		for _, line := range strings.Split(convert, "\n")[1:] {
			x := strings.Split(line, " ")
			var temp []int
			for _, y := range x {
				num, _ := strconv.Atoi(y)
				temp = append(temp, num)
			}
			numbers[i] = append(numbers[i], temp)
		}
	}

	return seeds, numbers
}

func convertThingToTheNextThing(number int, maps [][]int) int {
	for _, v := range maps {
		source := v[1]
		length := v[2]
		if number >= source && number < source+length {
			destination := v[0]
			source := v[1]
			return number + destination - source
		}
	}
	return number
}

func day5part2() int {
	seeds, numbers := parseInput5()
	location := math.MaxInt

	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		length := seeds[i+1]
		println(seed, length)
		for j := 0; j < length; j++ {
			thing := seed + j
			for _, v := range numbers {
				thing = convertThingToTheNextThing(thing, v)
			}
			if thing < location {
				location = thing
			}
		}
	}

	return location
}
