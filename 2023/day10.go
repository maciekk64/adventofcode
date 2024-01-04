package main

import (
	"math"
	"strings"
)

type pipe struct{ y, x, direction int }

var pipes = map[string]struct{ first, second int }{
	"|": {0, 2},
	"-": {1, 3},
	"L": {0, 1},
	"J": {0, 3},
	"7": {2, 3},
	"F": {1, 2},
}

func day10() int {
	x := readInput("in/day10.in")
	input := strings.Split(x, "\n")
	startIndex := strings.Index(x, "S")
	start := struct{ y, x int }{startIndex / (len(input) + 1), startIndex % (len(input) + 1)}

	next := pipe{start.y - 1, start.x, 2}
	returnValue := loopFarthestPoint(next, input)
	if returnValue != 0 {
		return returnValue
	}

	next = pipe{start.y, start.x + 1, 3}
	returnValue = loopFarthestPoint(next, input)
	if returnValue != 0 {
		return returnValue
	}

	next = pipe{start.y + 1, start.x, 0}
	returnValue = loopFarthestPoint(next, input)
	if returnValue != 0 {
		return returnValue
	}

	next = pipe{start.y, start.x - 1, 1}
	returnValue = loopFarthestPoint(next, input)
	if returnValue != 0 {
		return returnValue
	}

	return 0
}

func loopFarthestPoint(next pipe, input []string) int {
	for i := 1; next.direction != -1; i++ {
		next = nextPipe(input, next)
		if next.direction == -2 {
			return int(math.Ceil(float64(i) / 2))
		}
	}
	return 0
}

func nextPipe(input []string, curr pipe) pipe {
	currPipe := string(input[curr.y][curr.x])
	if currPipe == "S" {
		return pipe{0, 0, -2}
	}
	next := nextDirection(currPipe, curr.direction)
	switch next {
	case 0:
		return pipe{curr.y - 1, curr.x, 2}
	case 1:
		return pipe{curr.y, curr.x + 1, 3}
	case 2:
		return pipe{curr.y + 1, curr.x, 0}
	case 3:
		return pipe{curr.y, curr.x - 1, 1}
	}
	return pipe{0, 0, -1}
}

func nextDirection(pipeSymbol string, direction int) int {
	pipe := pipes[pipeSymbol]
	if pipe.first == direction {
		return pipe.second
	} else if pipe.second == direction {
		return pipe.first
	}
	return -1
}
