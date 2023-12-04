package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const LineBreak = "\r\n"

func day1(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day1/%s.txt", filename))
	if err != nil {
        panic(err)
	}

	switch part {
	case "1": day1Part1(content)
	case "2": day1Part2(content)
    default: panic(fmt.Sprintf("Unknown part: %s", part))
	}
}

func day1Part1(content []byte) {
	numbers := make([]int, 0)

	lines := strings.Split(string(content), LineBreak)
	for _, line := range lines {
		if line == "" {
			continue
		}

		digits := make([]string, 0)
		for _, char := range line {
			switch {
			case '0' <= char && char <= '9':
				digits = append(digits, string(char))
			default:
				continue
			}
		}

		n := 0
		switch len(digits) {
		case 0:
			continue
		default:
			var err error = nil
			n, err = strconv.Atoi(digits[0] + digits[len(digits)-1])
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
		}

		numbers = append(numbers, n)
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	fmt.Printf("Sum: %d\n", sum)
}

type Number struct {
	Value    int
	Position int
}

type Digit struct {
	Name  string
	Value int
}

func day1Part2(content []byte) {
	digits := []Digit{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
		{"0", 0},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
	}

	ns := make([]int, 0)
	lines := strings.Split(string(content), LineBreak)
	for _, line := range lines {
		numbers := make([]Number, 0)
		if line == "" {
			continue
		}

		for _, digit := range digits {
			index := strings.Index(line, digit.Name)
			if index != -1 {
				numbers = append(numbers, Number{digit.Value, index})
			}

			index = strings.LastIndex(line, digit.Name)
			if index != -1 {
				numbers = append(numbers, Number{digit.Value, index})
			}
		}

		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i].Position < numbers[j].Position
		})

		if len(numbers) == 0 {
			continue
		}

		n := 10*numbers[0].Value + numbers[len(numbers)-1].Value

		ns = append(ns, n)
	}

	sum := 0
	for _, n := range ns {
		sum += n
	}

	fmt.Printf("Sum: %d\n", sum)
}
