package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func day14(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day14/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)

    rocks := parseDay14(input)

    switch part {
    case "1":
        day14Part1(rocks)
    case "2":
        panic("not implemented")
    }
}

func day14Part1(rocks [][]rune) {
    tilt(rocks)
}

func tilt(rocks [][]rune) {
    rows := len(rocks)
    cols := len(rocks[0])

    load := 0
    for col := 0; col < cols; col++ {
        canMoveTo := 0
        for y, line := range rocks {
            c := line[col]

            switch c {
                case '#': canMoveTo = y + 1
                case 'O':
                    // NOTE (Emil): Move rock
                    if canMoveTo < y {
                        rocks[y][col] = '.'
                        rocks[canMoveTo][col] = 'O'

                        load += (rows - canMoveTo)
                        canMoveTo = canMoveTo + 1
                    } else {
                        load += (rows - y)
                        canMoveTo = y + 1
                    }
                case '.': continue
            }
        }
    }

    fmt.Printf("Load: %d\n", load)
}

func printRocks(rocks [][]rune) {
    for _, line := range rocks {
        fmt.Println(string(line))
    }
}

func parseDay14(input string) [][]rune {
    reader := strings.NewReader(input)
    scanner := bufio.NewScanner(reader)

    rocks := make([][]rune, 0)
    for scanner.Scan() {
        line := scanner.Text()

        r := []rune(line)

        rocks = append(rocks, r)
    }

    return rocks
}

