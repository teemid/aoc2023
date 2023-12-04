package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	GEAR = iota
	OTHER
)

type PartNumber struct {
	start int
	end   int
}

type Symbol struct {
	x int
	y int
	t int
}

type Part struct {
	x    int
	y    int
	part *PartNumber
}

func NewPart() *PartNumber {
	return &PartNumber{-1, 0}
}

func day3(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day3/%s.txt", filename))
	if err != nil {
        panic(err)
	}

    input := string(content)
    data := parseDay3Input(input)

    switch part {
    case "1": day3Part1(input, data); break
    case "2": day3Part2(input, data); break
    default: panic(fmt.Sprintf("Unknown part: %s", part))
    }

}

func day3Part1(input string, data *DataDay3) {
	validParts := make(map[*PartNumber]*PartNumber)

	for _, symbol := range data.symbols {
		for y := symbol.y - 1; y <= symbol.y+1; y++ {
			if y < 0 || y >= data.rowCount {
				continue
			}

			partRow := data.partPositions[y]
			for _, part := range partRow {
				diffX := abs(part.x - symbol.x)
				diffY := abs(part.y - symbol.y)
				if diffX <= 1 && diffY <= 1 {
					validParts[part.part] = part.part
				}
			}
		}
	}
    
	sum := 0
	for _, part := range validParts {
		n, err := strconv.Atoi(input[part.start : part.end+1])
		if err != nil {
			panic(err)
		}

		sum += n
	}

	fmt.Printf("Sum: %d\n", sum)
}

func day3Part2(input string, data *DataDay3) {
	ratioSum := 0
	for _, symbol := range data.symbols {
		adjacentParts := make(map[*PartNumber]*PartNumber)

		for y := symbol.y - 1; y <= symbol.y+1; y++ {
			if y < 0 || y >= data.rowCount {
				continue
			}

			partRow := data.partPositions[y]
			for _, part := range partRow {
				diffX := abs(part.x - symbol.x)
				diffY := abs(part.y - symbol.y)
				if diffX <= 1 && diffY <= 1 {
					adjacentParts[part.part] = part.part
				}
			}
		}

		if symbol.t == GEAR && len(adjacentParts) == 2 {
			parts := make([]*PartNumber, 0, 2)
			for _, part := range adjacentParts {
				parts = append(parts, part)
			}

			n1 := parts[0]
			n2 := parts[1]

			c1, err := strconv.Atoi(input[n1.start : n1.end+1])
			if err != nil {
                panic(err)
			}
			c2, err := strconv.Atoi(input[n2.start : n2.end+1])
			if err != nil {
                panic(err)
			}

			ratio := c1 * c2
			ratioSum += ratio
		}
	}

    fmt.Printf("Ratio sum: %d\n", ratioSum)
}

type DataDay3 struct {
    rowCount int
    symbols []Symbol
    partPositions [][]Part
    parts []*PartNumber
}

func parseDay3Input(input string) *DataDay3 {
    data := &DataDay3{
        rowCount: 0,
        symbols: make([]Symbol, 0),
        partPositions: make([][]Part, 0),
        parts: make([]*PartNumber, 0),
    }

	colCount := 0

	rowParts := make([]Part, 0)

	part := NewPart()

	row := 0
	col := -1
	n := -1

	for _, char := range input {
		n++
		col++

		data.rowCount = max(data.rowCount, row+1)

		switch char {
		case ' ':
			continue
		case '.':
			if part.start != -1 {
				data.parts = append(data.parts, part)
				part = NewPart()
			}
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if part.start == -1 {
				part.start = n
				part.end = n
			} else {
				part.end = n
			}

			rowParts = append(rowParts, Part{col, row, part})
		case '\r':
			continue
		case '\n':
			// Assumes part numbers don't cross lines
			if part.start != -1 {
				data.parts = append(data.parts, part)
				part = NewPart()
			}

			data.partPositions = append(data.partPositions, rowParts)
			rowParts = make([]Part, 0)

			colCount++

			col = -1
			row++
			break
		default:
			if part.start != -1 {
				data.parts = append(data.parts, part)
				part = NewPart()
			}

			t := OTHER
			if char == '*' {
				t = GEAR
			}

			data.symbols = append(data.symbols, Symbol{col, row, t})
		}
	}

    return data
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
