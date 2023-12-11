package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type GalaxyImage struct {
    galaxies []Point
    image [][]rune
}

func day11(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day11/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)

    switch part {
    case "1":
        positions := day11ParseInput(input, 1)
        day11Part1(positions)
    case "2":
        positions := day11ParseInput(input, 1_000_000-1)
        day11Part1(positions)
    }
}

type GalaxyPair struct {
    g1 Point
    g2 Point
}

func day11Part1(positions []Point64) {
    sum := int64(0)
    ps := positions[:]
    for len(ps) > 1 {
        p1 := ps[0]
        ps = ps[1:]
        for _, p2 := range ps {
            d := distance(p1, p2)
            
            sum += d
        }
    }

    fmt.Printf("sum: %d\n", sum)
}

func distance(p1, p2 Point64) int64 {
    return abs64(p1.x - p2.x) + abs64(p1.y - p2.y)
}

func abs64(i int64) int64 {
    if i < 0 {
        return -i
    }

    return i
}

type Point64 struct {
    x int64
    y int64
}

func day11ParseInput(input string, offset int64) []Point64 {
    scanner := bufio.NewScanner(strings.NewReader(input))
    scanner.Split(bufio.ScanLines)

    positions := make([]Point64, 0)
    rows := make([]int64, 0)
    cols := make([]bool, 0)

    m := make([][]rune, 0)

    yOffset := int64(0)
    for y := 0; scanner.Scan(); y++ {
        line := scanner.Text()
        blankLine := true
        for x, char := range line {
            if len(cols) == 0 {
                cols = make([]bool, len(line))
                for i := range cols {
                    cols[i] = true
                }
            }

            m = append(m, []rune(line))

            if char == '#' {
                positions = append(positions, Point64{int64(x), int64(y)})
                blankLine = false
                cols[x] = false
            }
        }

        if blankLine {
            yOffset++
        }

        rows = append(rows, yOffset)
    }

    xOffset := int64(0)
    columns := make([]int64, len(cols))
    for i, col := range cols {
        if col {
            xOffset++
        }

        columns[i] = xOffset
    }

    for i, pos := range positions {
        pos.x += columns[pos.x] * offset
        pos.y += rows[pos.y] * offset

        positions[i] = pos
    }

    return positions
}

