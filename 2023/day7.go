package main

import (
	"sort"
	"strconv"
	"strings"
)

func day7() (sum int) {
	input := parseInput7()
	sort.Slice(input, func(i, j int) bool {
		return less(input[i].Hand, input[j].Hand)
	})

	for i, v := range input {
		sum += (i + 1) * v.Bid
	}

	return
}

func less(first, second string) bool {
	sameFirst, sameSecond := count(first), count(second)
	length := min(len(sameFirst), len(sameSecond))

	for i := 0; i < length; i++ {
		if sameFirst[i] < sameSecond[i] {
			return true
		} else if sameFirst[i] > sameSecond[i] {
			return false
		}
	}

	return stronger(first, second)
}

var strength = map[string]int{
	"A": 2,
	"K": 3,
	"Q": 4,
	// "J": 5,
	"J": 15,
	"T": 6,
	"9": 7,
	"8": 8,
	"7": 9,
	"6": 10,
	"5": 11,
	"4": 12,
	"3": 13,
	"2": 14,
}

func stronger(first string, second string) bool {
	for i := 0; i < 5; i++ {
		if strength[string(first[i])] > strength[string(second[i])] {
			return true
		} else if strength[string(first[i])] < strength[string(second[i])] {
			return false
		}
	}
	panic("Should not be equal")
}

func count(cards string) []int {
	same := map[string]int{}
	for _, v := range cards {
		same[string(v)]++
	}

	if same["J"] > 0 {
		var maxCount string
		for i, v := range same {
			if v > same[maxCount] && i != "J" {
				maxCount = i
			}
		}
		same[maxCount] += same["J"]
		delete(same, "J")
	}

	sameCount := []int{}
	for _, v := range same {
		sameCount = append(sameCount, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sameCount)))

	return sameCount
}

func parseInput7() (input []struct {
	Hand string
	Bid  int
},
) {
	for _, v := range strings.Split(readInput("in/day7.in"), "\n") {
		values := strings.Split(v, " ")
		hand := values[0]
		bid, _ := strconv.Atoi(values[1])
		input = append(input, struct {
			Hand string
			Bid  int
		}{hand, bid})
	}
	return
}
