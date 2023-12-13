package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SpringRow struct {
    row []rune
    segments []int
}

func day12(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day12/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)
    rows := parseDay12(input)

    switch part {
    case "1":
        day12Part1(rows)
    case "2":
        day12Part2(rows)
    }
}

func day12Part1(rows []SpringRow) {
    sum := 0
    for _, row := range rows {
        rows := generate(row.row)

        printSpringRow(row.row)
        count := 0
        for _, r := range rows {
            valid := validateSprintRow(r, row.segments)
            if valid {
                count++
            }
        }

        sum += count
    }

    fmt.Printf("Sum: %d\n", sum)
}

func day12Part2(rows []SpringRow) {
    sum := 0

    m := 0
    m5 := 0
    for _, row := range rows {
        s := string(row.row)
        l := len(s)
        l5 := l * 5
        fmt.Printf("Row: %s, len: %d, len * 5: %d\n", s, l, l5)

        m = max(m, l)
        m5 = max(m5, l5)
        // rows := generate(row.row)


        // count := 0
        // for _, r := range rows {
        //     valid := validateSprintRow(r, row.segments)
        //     if valid {
        //         count++
        //     }
        // }

        // sum += count
    }

    fmt.Printf("Max: %d, Max * 5: %d\n", m, m5)

    fmt.Printf("Sum: %d\n", sum)
}

func unfold(row []rune, segments []int) ([]rune, []int) {
    s := string(row)
    r := strings.Join([]string{s, s, s, s, s}, "?")

    segs := make([]int, 0) 
    for i := 0; i < 5; i++ {
        segs = append(segs, segments...)
    }

    return []rune(r), segs
}

func printSpringRow(row []rune) {
    for _, c := range row {
        fmt.Printf("%c", c)
    }
    fmt.Printf("\n")
}

func validateSprintRow(row []rune, segments []int) bool {
    currentRange := 0
    segmentIndex := 0
    segment := segments[segmentIndex]
    for _, c := range row {
        switch c {
        case '.':
            if currentRange > 0 {
                if currentRange != segment {
                    return false
                }

                currentRange = 0
                segmentIndex++
                if segmentIndex < len(segments) {
                    segment = segments[segmentIndex]
                }
            }
        case '#': currentRange++
        }
    }

    if currentRange > 0 {
        if currentRange != segment {
            return false
        }

        currentRange = 0
        segmentIndex++
        if segmentIndex < len(segments) {
            segment = segments[segmentIndex]
        }
    }

    return segmentIndex == len(segments)
}

func generate(row []rune) [][]rune {
    rows := make([][]rune, 0)
    r := make([]rune, len(row))
    copy(r, row)
    rows = append(rows, r)
    for i, c := range row {
        switch c {
        case '?':
            for _, r := range rows {
                a := make([]rune, len(r))
                copy(a, r)
                r[i] = '.'
                a[i] = '#'
                rows = append(rows, a)
            }
        }
    }

    return rows
}

func parseDay12(input string) []SpringRow {
    rows := make([]SpringRow, 0)

    scanner := bufio.NewScanner(strings.NewReader(input))
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        line := scanner.Text()

        parts := strings.Split(line, " ")
        row := SpringRow{
            row: []rune(parts[0]),
            segments: make([]int, 0),
        }

        segments := strings.Split(parts[1], ",")
        for _, segment := range segments {
            n, err := strconv.Atoi(segment)
            if err != nil {
                panic(err)
            }

            row.segments = append(row.segments, n)
        }

        rows = append(rows, row)
    }

    return rows
}
