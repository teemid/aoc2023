package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func day6(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day6/%s.txt", filename))
	if err != nil {
		panic(err)
	}

	input := string(content)

	switch part {
	case "1":
		races := parseDay6Input(input)
		answer := solveRaces(races)

		fmt.Printf("Day 6 part 1: %v\n", answer)
	case "2":
		input = strings.ReplaceAll(input, " ", "")
		races := parseDay6Input(input)
		answer := solveRaces(races)

		fmt.Printf("Day 6 part 1: %v\n", answer)
	}
}

func solveRaces(races []Race) int {
	prod := 1

	for _, race := range races {
		x1, x2 := solve(float64(race.Time), float64(-race.Distance-1))

		ways := abs(x1-x2) + 1

		prod = prod * ways
	}

	return prod
}

func solve(b float64, c float64) (int, int) {
	inner := math.Sqrt(math.Pow(b, 2) - (4.0 * -1.0 * c))
	denom := 2.0 * -1.0

	x1 := (-b + inner) / denom
	x2 := (-b - inner) / denom

	return int(math.Ceil(x1)), int(math.Floor(x2))
}

func parseDay6Input(input string) []Race {
	races := make([]Race, 0)
	idx := 0

	race := Race{}

	start := -1
	end := 0
	isTime := true

	for i, char := range input {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if start == -1 {
				start = i
				end = i
			} else {
				end = i
			}
		case ' ':
			if start != -1 {
				n, err := strconv.Atoi(input[start : end+1])
				if err != nil {
					panic(err)
				}

				if isTime {
					race.Time = n
					races = append(races, race)
				} else {
					races[idx].Distance = n
					idx++
				}

				start = -1
				end = 0
			}
		case '\r':
		case '\n':
			n, err := strconv.Atoi(input[start : end+1])
			if err != nil {
				panic(err)
			}

			if isTime {
				race.Time = n
				races = append(races, race)
			} else {
				races[idx].Distance = n
				idx++
			}

			isTime = false
			idx = 0
			start = -1
			end = 0
		default:
			continue
		}
	}

	return races
}
