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
    end int
}

type Symbol struct {
    x int
    y int
    t int
}

type Part struct {
    x int 
    y int
    part *PartNumber
}

func NewPart() *PartNumber {
    return &PartNumber{-1, 0}
}

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        fmt.Println("Usage: main <filename>")
        os.Exit(1)
    }

    content, err := os.ReadFile(args[0])
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    strconv.Atoi("1")

    symbols := make([]Symbol, 0)
    partPositions := make([][]Part, 0)
    parts := make([]*PartNumber, 0)

    part := NewPart()

    input := string(content)

    rowCount := 0
    colCount := 0

    rowParts := make([]Part, 0)

    row := 0
    col := -1
    n := -1
    for _, char := range input {
        n++
        col++

        rowCount = max(rowCount, row + 1)

        switch char {
        case ' ': continue
        case '.': 
            if part.start != -1 {
                parts = append(parts, part)
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
        case '\r': continue
        case '\n': 
            // Assumes part numbers don't cross lines
            if part.start != -1 {
                parts = append(parts, part)
                part = NewPart()
            }

            partPositions = append(partPositions, rowParts)
            rowParts = make([]Part, 0)
            
            colCount++

            col = -1
            row++
            break
        default:
            if part.start != -1 {
                parts = append(parts, part)
                part = NewPart()
            }

            t := OTHER
            if char == '*' {
                t = GEAR
            }

            symbols = append(symbols, Symbol{col, row, t})
        }
    }

    fmt.Printf("Rows count: %d\n", rowCount)
    fmt.Printf("Column count: %d\n", colCount)

    validParts := make(map[*PartNumber]*PartNumber)

    ratioSum := 0
    for _, symbol := range symbols {
        adjacentParts := make(map[*PartNumber]*PartNumber)

        for y := symbol.y - 1; y <= symbol.y + 1; y++ {
            if y < 0 || y >= rowCount {
                continue
            }

            partRow := partPositions[y];
            for _, part := range partRow {
                diffX := abs(part.x - symbol.x)
                diffY := abs(part.y - symbol.y)
                if diffX <= 1 && diffY <= 1 {
                    validParts[part.part] = part.part 
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

            c1, err := strconv.Atoi(input[n1.start:n1.end+1])
            if err != nil {
                fmt.Printf("Error: %s\n", err)
                os.Exit(1)
            }
            c2, err := strconv.Atoi(input[n2.start:n2.end+1])
            if err != nil {
                fmt.Printf("Error: %s\n", err)
                os.Exit(1)
            }

            ratio := c1 * c2
            ratioSum += ratio
        }
    }

    sum := 0
    for _, part := range validParts {
        n, err := strconv.Atoi(input[part.start:part.end+1])
        if err != nil {
            panic(err)
        }

        sum += n
    }

    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Ratio sum: %d\n", ratioSum)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}
