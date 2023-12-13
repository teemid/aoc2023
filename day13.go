package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AshMap struct {
    rows []uint
    cols []uint
}

func day13(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day13/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)
    m := day13Parse(input)

    switch part {
    case "1":
        day13Part1(m)
    case "2":
        panic("Not implemented yet")
    }
}

func day13Part1(m []AshMap) int {
    for _, m := range m {
        fmt.Printf("Rows:\n")
        for _, r := range m.rows {
            fmt.Printf(fmt.Sprintf("%%0%db\n", len(m.cols)), r)
        }

        fmt.Println()

        fmt.Printf("Cols:\n")
        for _, c := range m.cols {
            fmt.Printf(fmt.Sprintf("%%0%db\n", len(m.rows)), c)
        }

        fmt.Println()
    }

    return 0
}

type Dimensions struct {
    width int
    height int
}

func day13Parse(input string) []AshMap {
    ms := make([]AshMap, 0)
    reader := strings.NewReader(input)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanLines)

    maps := make([][]string, 0)
    lines := make([]string, 0)
    dimensions := make([]Dimensions, 0)
    d := Dimensions{
        width: 0,
        height: 0,
    }

    for scanner.Scan() {
        text := scanner.Text()
        if d.width == 0 {
            d.width = len(text)
        }

        fmt.Printf("text: %v\n", text)
        if text == "\r\n" || text == "\n" || text == "" {
            dimensions = append(dimensions, d)
            d = Dimensions{
                width: 0,
                height: 0,
            }

            maps = append(maps, lines)
            lines = make([]string, 0)
        } else {
            d.height++
        }

        lines = append(lines, text)
    }

    maps = append(maps, lines)
    dimensions = append(dimensions, d)

    fmt.Printf("maps: %v\n", maps)

    fmt.Printf("dimensions: %v\n", dimensions)
    for i, Map := range maps {
        d := dimensions[i]
        m := AshMap{
            rows: make([]uint, d.height),
            cols: make([]uint, d.width),
        }

        for x, line := range Map {
            if len(m.cols) == 0 {
                m.cols = make([]uint, d.width)
            }

            fmt.Printf("line: %v\n", line)

            l := uint(0)
            for j, c := range line {
                if c == '#' {
                    fmt.Printf("width: %v, j: %v\n", d.width, j)
                    fmt.Printf("height: %v, x: %v\n", d.height, x)
                    l |= 1 << (d.width - j - 1)
                    m.cols[j] |= 1 << (d.height - x - 1)
                }
            }

            m.rows[x] = l
        }

        ms = append(ms, m)
    }

    return ms
}

