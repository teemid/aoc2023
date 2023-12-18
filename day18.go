package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DigInstruction struct {
    direction string
    distance int
    color string
}

type Position struct {
    x int
    y int
}

func day18(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day18/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)

    instructions := parseDay18Input(input)
    switch part {
    case "1":
        day18Part1(instructions)
    }
}

func day18Part1(instructions []DigInstruction) {
    positions := make([]Position, 0)

    minPos := Position{0, 0}
    maxPos := Position{0, 0}
    pos := Position{0, 0}
    for _, instruction := range instructions {
        dir := Position{0, 0}
        switch instruction.direction {
            case "U":
                dir = Position{0, -1}
            case "R":
                dir = Position{1, 0}
            case "D":
                dir = Position{0, 1}
            case "L":
                dir = Position{-1, 0}
        }

        for i := 0; i < instruction.distance; i++ {
            new := Position{pos.x + dir.x, pos.y + dir.y}
            pos = new

            if new.x <= minPos.x {
                minPos.x = new.x
            }

            if new.y <= minPos.y {
                minPos.y = new.y
            }

            if new.x >= maxPos.x {
                maxPos.x = new.x
            }

            if new.y >= maxPos.y {
                maxPos.y = new.y
            }

            positions = append(positions, new)
        }
    }

    offset := Position{-minPos.x, -minPos.y}

    fmt.Printf("Min: %v\n", minPos)
    fmt.Printf("Max: %v\n", maxPos)
    fmt.Printf("Offset: %v\n", offset)

    m := make([][]rune, maxPos.y + offset.y + 2)
    for i := range m {
        m[i] = make([]rune, maxPos.x + offset.x + 2)

        for j := range m[i] {
            m[i][j] = '.'
        }
    }

    count := 0
    for _, pos := range positions {
        m[pos.y + offset.y][pos.x + offset.x] = '#'

        count++
    }

    printDigPlan(m)

    for y, row := range m {
        c := 0
        for x := 0; x < len(row); {
            col := row[x]
            
            switch col {
            case '#':
                l := 1
                x++
                col = row[x]
                for col == '#' {
                    fmt.Printf("%c", col)
                    fmt.Printf("%d", c)

                    l++ 
                    x++
                    col = row[x]
                }

                if l == 1 {
                    c++
                }
            case '.':
                if c % 2 == 1 {
                    m[y][x] = '#'
                }

                x++

                fmt.Printf("%c", col)
                fmt.Printf("%d", c)
            }
        }

        fmt.Println()
    }

    printDigPlan(m)

    fmt.Printf("Count: %d\n", count)
}

func calculateDistances(positions []Position, offset Position, maxPos Position) int {
    lines := make([][]Position, maxPos.y + offset.y + 1)
    for _, pos := range positions { 
        lines[pos.y + offset.y] = append(lines[pos.y + offset.y], pos)
    }

    count := 0
    for _, line := range lines {
        sort.Slice(line, func(i, j int) bool {
            return line[i].x < line[j].x
        })

        fmt.Printf("%v, len = %d\n", line, len(line))

        p := line[0]
        for _, pos := range line { 
            d := p.distance(pos)
            if d > 1 {
                count += d
            } else {
                count++
            }

            p = pos
        }
    }

    return count
}

func (p *Position) distance(other Position) int {
    return abs(p.x - other.x) + abs(p.y - other.y)
}

func printDigPlan(plan [][]rune) {
    for _, row := range plan {
        for _, col := range row {
            fmt.Printf("%c", col)
        }
        fmt.Printf("\n")
    }
}

func parseDay18Input(input string) []DigInstruction {
    reader := strings.NewReader(input)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanLines)

    instructions := make([]DigInstruction, 0)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")
        if len(parts) != 3 {
            panic("Invalid input")
        }

        distance, err := strconv.Atoi(parts[1])
        if err != nil {

        }

        instruction := DigInstruction{
            direction: parts[0],
            distance: distance,
            color: parts[2][1:len(parts[2])-1],
        }

        instructions = append(instructions, instruction)
    }

    return instructions
}
