package day1

import (
	utils "aoc_21/utilities"
	"log"
	"strconv"
)

func RunDay1() {
	log.Println("Hello Day1")

	// Open input file
	fileH := utils.ReadFile("day1/input.txt")

	// Store each line in array
	lines := utils.ReadLinesIntoStringArr(fileH)

	part1(lines)

	part2(lines)

}

func part1(lines []string) {
	if len(lines) != 0 {
		var currentNum int
		var nextNum int
		var incrementedNumbers int64
		for idx, item := range lines {
			// Get current Number
			currentNum, _ = strconv.Atoi(item)

			// Get Next number as long as its not out of index
			if len(lines) > idx + 1 {
				nextNum,_ = strconv.Atoi(lines[idx+1])
				if nextNum > currentNum {
					incrementedNumbers += 1
				}
			}
		}
		log.Println("Part 1 Number of increments: ", strconv.FormatInt(incrementedNumbers, 10))
	}
}

func part2(lines []string) {
	if len(lines) != 0 {
		var currentNum int
		var nextNum int
		var nextNumPlus1 int
		var intArr []int
		var incrementedNumbers int64
		for idx, item := range lines {
			// Get current Number
			currentNum, _ = strconv.Atoi(item)

			// Get Next number as long as its not out of index
			if len(lines) > idx + 2 {
				nextNum,_ = strconv.Atoi(lines[idx+1])
				nextNumPlus1,_ = strconv.Atoi(lines[idx+2])
				intArr = append(intArr, currentNum+nextNum+nextNumPlus1)
			}
		}

		for idx, item := range intArr {
			currentNum = item
			if len(intArr) > idx + 1 {
				nextNum = intArr[idx+1]
				if nextNum > currentNum {
					incrementedNumbers += 1
				}
			}
		}
		log.Println("Part 2 Number of increments: ", strconv.FormatInt(incrementedNumbers, 10))
	}
}
