package main

import (
	"strconv"
	"strings"
)

func day1() (sum int) {
	input := strings.Split(readInput("in/day1.in"), "\n")
	for _, v := range input {
		num := ""
		temp := ""
		letters := ""
		for _, i := range v {
			if i > 48 && i <= 57 {
				temp = string(i)
				letters = ""
			} else {
				letters += string(i)
				for j, l := 0, len(letters)-2; j < l; j++ {
					switch letters[j:] {
					case "one":
						temp = "1"
					case "two":
						temp = "2"
					case "three":
						temp = "3"
					case "four":
						temp = "4"
					case "five":
						temp = "5"
					case "six":
						temp = "6"
					case "seven":
						temp = "7"
					case "eight":
						temp = "8"
					case "nine":
						temp = "9"
					}
				}
			}
			if num == "" {
				num = temp
			}
		}
		add, _ := strconv.Atoi(num + temp)
		sum += add
	}
	return
}
