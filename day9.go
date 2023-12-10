package main

import (
	"fmt"
	"os"
	"strconv"
)

func day9(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day9/%s.txt", filename))
	if err != nil {
		panic(err)
	}

	input := string(content)
	report := parseDay9Input(input)

	switch part {
	case "1":
		day9Part1(report)
	case "2":
		day9Part2(report)
	}
}

func day9Part1(report [][]int) {
	sum := 0
	for _, history := range report {
		value := predict(history)

		sum += value
	}

	fmt.Printf("Sum: %d\n", sum)
}

func day9Part2(report [][]int) {
	sum := 0
	for _, history := range report {
		value := extrapolateBackwards(history)

		sum += value
	}

	fmt.Printf("Sum: %d\n", sum)
}

func extrapolateBackwards(history []int) int {
	curr := history

	diffs := make([][]int, 0)
	diffs = append(diffs, history)
	for !isAllZero(curr) {
		diff := difference(curr)

		diffs = append(diffs, diff)
		curr = diff
	}

	last := 0
	for i := range diffs {
		d := diffs[len(diffs)-1-i]
		last = d[0] - last
		d = append([]int{last}, d...)
		diffs[len(diffs)-1-i] = d
	}

	return diffs[0][0]
}

func predict(history []int) int {
	curr := history

	diffs := make([][]int, 0)
	diffs = append(diffs, history)
	for !isAllZero(curr) {
		diff := difference(curr)

		diffs = append(diffs, diff)
		curr = diff
	}

	last := 0
	for i := range diffs {
		d := diffs[len(diffs)-1-i]
		last = d[len(d)-1] + last
		d = append(d, last)
		diffs[len(diffs)-1-i] = d
	}

	return diffs[0][len(diffs[0])-1]
}

func difference(series []int) []int {
	diff := make([]int, len(series)-1)
	for i := 0; i < len(series)-1; i++ {
		diff[i] = series[i+1] - series[i]
	}

	return diff
}

func isAllZero(diff []int) bool {
	for _, d := range diff {
		if d != 0 {
			return false
		}
	}

	return true
}

func parseDay9Input(input string) [][]int {
	report := make([][]int, 0)
	history := make([]int, 0)
	for i := 0; i < len(input); {
		c := input[i]

		switch c {
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n, j := parseNumber(i, input)
			if j == -1 {
				panic(fmt.Sprintf("Unexpected end of input: %s", input))
			}

			i = j

			history = append(history, n)
		case ' ':
			i++
		case '\r':
			i++
		case '\n':
			report = append(report, history)
			history = make([]int, 0)

			i++
		default:
			panic(fmt.Sprintf("Unexpected character: %c", c))
		}
	}

	return report
}

func parseNumber(i int, input string) (int, int) {
	j := 0
	for {
		c := input[i+j]
		switch c {
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			j++
		default:
			n, err := strconv.Atoi(input[i : i+j])
			if err != nil {
				panic(err)
			}

			return n, i + j
		}
	}
}
