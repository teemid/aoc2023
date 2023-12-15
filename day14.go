package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "time"
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
        printRocks(rocks)
        fmt.Println()
        start := time.Now()
        for i := 0; i < 1000000000; i++ {
            tilt(rocks, "north")
            tilt(rocks, "west")
            tilt(rocks, "south")
            tilt(rocks, "east")

            // printRocks(rocks)
            // fmt.Println()
            // fmt.Printf("Iteration: %d\n", i)
            // fmt.Printf("Time: %s\n", time.Since(start))
            if i % 100000 == 0 {
                fmt.Printf("Iteration: %d\n", i)
                fmt.Printf("Time: %s\n", time.Since(start))
            }
        }

        load := calculateLoad(rocks)
        fmt.Printf("Load: %d\n", load)
    }
}

func day14Part1(rocks [][]rune) {
    tilt(rocks, "north")
    load := calculateLoad(rocks)

    fmt.Printf("Load: %d\n", load)
}

func tiltAndLoad(rocks [][]rune) {
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

func tilt(rocks [][]rune, direction string) {
    switch direction {
    case "north": tiltNorth(rocks)
    case "south": tiltSouth(rocks)
    case "east": tiltEast(rocks)
    case "west": tiltWest(rocks)
    }
}

func tiltNorth(rocks [][]rune) {
    cols := len(rocks[0])

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

                        canMoveTo = canMoveTo + 1
                    } else {
                        canMoveTo = y + 1
                    }
                case '.': continue
            }
        }
    }
}

func tiltSouth(rocks [][]rune) {
    cols := len(rocks[0])

    for col := 0; col < cols; col++ {
        l := len(rocks)
        canMoveTo := l - 1
        for j := range rocks {
            line := rocks[l - j - 1]
            c := line[col]
            y := l - j - 1

            switch c {
                case '#': canMoveTo = y - 1
                case 'O':
                    // NOTE (Emil): Move rock
                    if canMoveTo > y {
                        rocks[y][col] = '.'
                        rocks[canMoveTo][col] = 'O'

                        canMoveTo = canMoveTo - 1
                    } else {
                        canMoveTo = y - 1
                    }
                case '.': continue
            }
        }
    }
}

func tiltEast(rocks [][]rune) {
    for y, line := range rocks {
        l := len(line)
        canMoveTo := l - 1
        for i := range line {
            x := l - i - 1
            c := line[x]

            switch c {
                case '#': canMoveTo = x - 1
                case 'O':
                    // NOTE (Emil): Move rock
                    if canMoveTo > x {
                        rocks[y][x] = '.'
                        rocks[y][canMoveTo] = 'O'

                        canMoveTo = canMoveTo - 1
                    } else {
                        canMoveTo = x - 1
                    }
                case '.': continue
            }
        }
    }
}

func tiltWest(rocks [][]rune) {
    for y, line := range rocks {
        canMoveTo := 0
        for x, c := range line {
            switch c {
                case '#': canMoveTo = x + 1
                case 'O':
                    // NOTE (Emil): Move rock
                    if canMoveTo < x {
                        rocks[y][x] = '.'
                        rocks[y][canMoveTo] = 'O'

                        canMoveTo = canMoveTo + 1
                    } else {
                        canMoveTo = x + 1
                    }
                case '.': continue
            }
        }
    }
}

func calculateLoad(rocks [][]rune) int {
    load := 0 

    rows := len(rocks)
    for y, line := range rocks {
        for _, c := range line {
            if c == 'O' {
                load += (rows - y)
            }
        }
    }

    return load
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

