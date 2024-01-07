package main

import (
	"strings"
)

type galaxy struct{ y, x int }

func day11() (sum int) {
	rows, cols, galaxies := parseInput11()
	for i, v := range galaxies {
		for _, w := range galaxies[i+1:] {
			sum += (galaxies_distance(v, w, rows, cols))
		}
	}
	return
}

func parseInput11() ([]int, []int, []galaxy) {
	input := strings.Split(readInput("in/day11.in"), "\n")
	lenght := len(input)
	empty_rows, empty_cols, galaxies := []int{}, []int{}, []galaxy{}
	for i := 0; i < lenght; i++ {
		empty_row, empty_col := true, true
		for j := 0; j < lenght; j++ {
			if string(input[i][j]) != "." {
				empty_row = false
				galaxies = append(galaxies, galaxy{i, j})
			}
			if string(input[j][i]) != "." {
				empty_col = false
			}
		}
		if empty_row {
			empty_rows = append(empty_rows, i)
		}
		if empty_col {
			empty_cols = append(empty_cols, i)
		}
	}
	return empty_rows, empty_cols, galaxies
}

func galaxies_distance(a, b galaxy, rows, cols []int) int {
	x1, x2, y1, y2 := max(a.x, b.x), min(a.x, b.x), max(a.y, b.y), min(a.y, b.y)
	expanded := 0
	for _, v := range cols {
		if v > x2 && v < x1 {
			expanded++
		}
	}
	for _, v := range rows {
		if v > y2 && v < y1 {
			expanded++
		}
	}
	// return x1 - x2 + y1 - y2 + expanded
	return x1 - x2 + y1 - y2 + expanded*999999
}
