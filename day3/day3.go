package day3

import (
	utils "aoc_21/utilities"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func RunDay3() {
	log.Println("Hello Day3")

	// Open input file
	fileH := utils.ReadFile("day3/input.txt")

	// Store each line in array
	lines := utils.ReadLinesIntoStringArr(fileH)

	part1(lines)
	part2(lines)
}

func part1(lines []string) {

	var holdingArr [][]string
	for _, item := range lines {
		tmp := strings.Split(item, "")
		holdingArr = append(holdingArr, tmp)
	}

	zeroCount := 0
	oneCount := 0
	gammaStr := ""
	epsilonStr := ""

	for i := 0; i < len(holdingArr[0]); i++ {
		for _, item := range holdingArr {
			if item[i] == "0" {
				zeroCount ++
			} else {
				oneCount ++
			}
		}

		if oneCount > zeroCount {
			gammaStr = gammaStr + "1"
			epsilonStr = epsilonStr + "0"
		} else {
			gammaStr = gammaStr + "0"
			epsilonStr = epsilonStr + "1"
		}

		zeroCount = 0
		oneCount = 0
	}

	gammaRateInt := convertStringBinaryToInt(gammaStr)
	epsilonRateInt := convertStringBinaryToInt(epsilonStr)

	powerConsumption := gammaRateInt * epsilonRateInt
	log.Println("Power Consumption: ", strconv.FormatInt(powerConsumption, 10))
}

func part2(lines []string) {

	var holdingArr [][]string
	for _, item := range lines {
		tmp := strings.Split(item, "")
		holdingArr = append(holdingArr, tmp)
	}

	oxygenSensorStr := reduceArray(holdingArr, "oxygen")
	co2SensorStr := reduceArray(holdingArr, "co2")
	oxygenSensorInt := convertStringBinaryToInt(strings.Join(oxygenSensorStr, ""))
	co2SensorInt := convertStringBinaryToInt(strings.Join(co2SensorStr, ""))

	oxyGenRating := oxygenSensorInt * co2SensorInt
	log.Println("Oxygen Generator Rating: ", strconv.FormatInt(oxyGenRating, 10))
}

func reduceArray(holdingArr [][]string, sensorType string) []string {

	zeroCount := 0
	oneCount := 0

	val1 := ""
	val2 := ""
	 if sensorType == "oxygen" {
		 val1 = "1"
		 val2 = "0"
	 } else {
		 val1 = "0"
		 val2 = "1"
	 }
	for i := 0; i < len(holdingArr[0]); i++ {
		for _, item := range holdingArr {
			if item[i] == "0" {
				zeroCount ++
			} else {
				oneCount ++
			}
		}

		if oneCount > zeroCount || oneCount == zeroCount {
			var tmpArr [][]string
			for _, item := range holdingArr {
				if item[i] == val1 {
					tmpArr = append(tmpArr, item)
				}
			}
			holdingArr = tmpArr

		} else if zeroCount > oneCount {
			var tmpArr [][]string
			for _, item := range holdingArr {
				if item[i] == val2 {
					tmpArr = append(tmpArr, item)
				}
			}
			holdingArr = tmpArr
		}

		if len(holdingArr) == 1 {
			return holdingArr[0]
		}

		zeroCount = 0
		oneCount = 0
	}

	return nil
}

func indexOf(element []string, data [][]string) (int) {
	for k, v := range data {
		if isArraysEqual(element, v) {
			return k
		}
	}
	return -1    //not found.
}

func isArraysEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func convertStringBinaryToInt (binaryNumStr string) int64 {

	if i, err := strconv.ParseInt(binaryNumStr, 2, 64); err != nil {
		fmt.Println(err)
		return -1
	} else {
		return i
	}
}