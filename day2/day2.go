package day2

import (
	utils "aoc_21/utilities"
	"log"
	"strconv"
	"strings"
)

type directionPairs struct {
	direction string
	movementCount int
}
func RunDay2() {
	log.Println("Hello Day2")

	// Open input file
	fileH := utils.ReadFile("day2/input.txt")

	// Store each line in array
	lines := utils.ReadLinesIntoStringArr(fileH)

	part1(lines)

	part2(lines)
}

func part1(lines []string) {

	// Create a map list of direction to value from each line
	dirMap := make([]directionPairs, 0)

	for _, item := range lines {
		tmp := strings.Split(item, " ")
		tmpNum, _ := strconv.Atoi(tmp[1])
		dirMap = append(dirMap, directionPairs{direction: tmp[0], movementCount: tmpNum})
	}

	horizPos := 0
	vertPos := 0

	for _, value := range dirMap {
		log.Println("Key: ", value.direction, "\tValue: ", strconv.FormatInt(int64(value.movementCount), 10))

		if value.direction == "forward" {
			horizPos = modifyNum(horizPos, value)
		} else {
			vertPos = modifyNum(vertPos, value)
		}
	}

	log.Println("Total value : ", strconv.FormatInt(int64(vertPos * horizPos), 10))
}

func modifyNum(currentVal int, pairVal directionPairs) int {

	switch pairVal.direction  {
	case "down":
		return currentVal + pairVal.movementCount
	case "forward":
		return currentVal + pairVal.movementCount
	case "up":
		return currentVal - pairVal.movementCount
	default:
		log.Println("UNKNOWN DIRECTION, doing nothing")
		return 0
	}
	return 0
}

func part2(lines []string) {
	// Create a map list of direction to value from each line
	dirMap := make([]directionPairs, 0)

	for _, item := range lines {
		tmp := strings.Split(item, " ")
		tmpNum, _ := strconv.Atoi(tmp[1])
		dirMap = append(dirMap, directionPairs{direction: tmp[0], movementCount: tmpNum})
	}

	horizPos := 0
	depth := 0
	aim := 0

	for _, value := range dirMap {
		log.Println("Key: ", value.direction, "\tValue: ", strconv.FormatInt(int64(value.movementCount), 10))

		if value.direction == "forward" {
			horizPos = horizPos + value.movementCount
			depth = depth + (aim * value.movementCount)
		} else if value.direction == "up" {
			aim = aim - value.movementCount
		} else if value.direction == "down" {
			aim = aim + value.movementCount
		}
	}

	log.Println("2nd part Total value : ", strconv.FormatInt(int64(depth * horizPos), 10))
}