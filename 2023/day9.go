package main

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput9() [][]int {
	input := strings.Split(readInput("in/day9.in"), "\n")
	all := [][]int{}
	for _, v := range input {
		history := []int{}
		for _, w := range strings.Split(v, " ") {
			num, _ := strconv.Atoi(w)
			history = append(history, num)
		}
		all = append(all, history)
	}
	return all
}

func day9part1() (sum int) {
	input := parseInput9()
	for _, v := range input {
		sum += predictNext(v)
	}
	return
}

func predictNext(history []int) (sum int) {
	for length := len(history) - 1; !isZeroed(history[:length]); length-- {
		for i := 0; i < length; i++ {
			history[i] = history[i+1] - history[i]
		}
	}

	for _, v := range history {
		sum += v
	}

	return
}

func isZeroed(history []int) bool {
	for _, v := range history {
		if v != 0 {
			return false
		}
	}
	return true
}

func day9part2() (sum int) {
	input := parseInput9()
	for _, v := range input {
		sum += predictPrev(v)
	}
	return
}

func predictPrev(history []int) (sum int) {
	slices.Reverse(history)
	for length := len(history) - 1; !isZeroed(history[:length]); length-- {
		for i := 0; i < length; i++ {
			history[i] = history[i] - history[i+1]
		}
	}

	for _, v := range history {
		sum = v - sum
	}

	return
}
