package main

import (
	"strconv"
	"strings"
)

func day3part1() int {
	input := strings.Split(readInput("in/day3.in"), "\n")
	temp := ""
	adjecent := false
	sum := 0
	width := len(input)
	for i, v := range input {
		for j := range v {
			if input[i][j] >= 48 && input[i][j] <= 57 {
				temp += string(input[i][j])
				if isAdjecent(i, j, width, input) {
					adjecent = true
				}
			} else {
				if adjecent {
					add, _ := strconv.Atoi(temp)
					sum += add
				}
				temp = ""
				adjecent = false
			}
		}
		if adjecent {
			add, _ := strconv.Atoi(temp)
			sum += add
		}
		temp = ""
		adjecent = false
	}
	return sum
}

func isAdjecent(i, j, width int, input []string) bool {
	if j+1 < width {
		if isSymbol(input[i][j+1]) {
			return true
		}
	}
	if j-1 >= 0 {
		if isSymbol(input[i][j-1]) {
			return true
		}
	}
	if i+1 < width {
		if isSymbol(input[i+1][j]) {
			return true
		}
		if j+1 < width {
			if isSymbol(input[i+1][j+1]) {
				return true
			}
		}
		if j-1 >= 0 {
			if isSymbol(input[i+1][j-1]) {
				return true
			}
		}
	}
	if i-1 >= 0 {
		if isSymbol(input[i-1][j]) {
			return true
		}

		if j+1 < width {
			if isSymbol(input[i-1][j+1]) {
				return true
			}
		}
		if j-1 >= 0 {
			if isSymbol(input[i-1][j-1]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(char byte) bool {
	return char != 46 && (char < 48 || char > 57)
}

func day3part2() int {
	input := strings.Split(readInput("in/day3.in"), "\n")
	temp := ""
	adjecent := [2]int{}
	gears := map[[2]int]int{}
	numbers := map[[2]int][]int{}
	sum := 0
	width := len(input)
	for i, v := range input {
		for j := range v {
			if input[i][j] >= 48 && input[i][j] <= 57 {
				temp += string(input[i][j])
				if isAdjecentToGear(i, j, width, input) != [2]int{0, 0} {
					adjecent = isAdjecentToGear(i, j, width, input)
				}
			} else {
				if adjecent != [2]int{0, 0} {
					gears[adjecent]++
					add, _ := strconv.Atoi(temp)
					numbers[adjecent] = append(numbers[adjecent], add)
				}
				temp = ""
				adjecent = [2]int{}
			}
		}
		if adjecent != [2]int{0, 0} {
			gears[adjecent]++
			add, _ := strconv.Atoi(temp)
			numbers[adjecent] = append(numbers[adjecent], add)
		}
		temp = ""
		adjecent = [2]int{}
	}
	for i := range gears {
		if len(numbers[i]) == 2 {
			sum += numbers[i][0] * numbers[i][1]
		}
	}
	return sum
}

func isAdjecentToGear(i, j, width int, input []string) [2]int {
	if j+1 < width {
		if isGear(input[i][j+1]) {
			return [2]int{i, j + 1}
		}
	}
	if j-1 >= 0 {
		if isGear(input[i][j-1]) {
			return [2]int{i, j - 1}
		}
	}
	if i+1 < width {
		if isGear(input[i+1][j]) {
			return [2]int{i + 1, j}
		}
		if j+1 < width {
			if isGear(input[i+1][j+1]) {
				return [2]int{i + 1, j + 1}
			}
		}
		if j-1 >= 0 {
			if isGear(input[i+1][j-1]) {
				return [2]int{i + 1, j - 1}
			}
		}
	}
	if i-1 >= 0 {
		if isGear(input[i-1][j]) {
			return [2]int{i - 1, j}
		}

		if j+1 < width {
			if isGear(input[i-1][j+1]) {
				return [2]int{i - 1, j + 1}
			}
		}
		if j-1 >= 0 {
			if isGear(input[i-1][j-1]) {
				return [2]int{i - 1, j - 1}
			}
		}
	}
	return [2]int{}
}

func isGear(char byte) bool {
	return string(char) == "*"
}
