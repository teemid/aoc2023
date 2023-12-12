package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SpringRow struct {
    row string
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
        fmt.Printf("Rows: %v\n", rows)
    case "2":
        panic("Not implemented yet")
    }
}

func parseDay12(input string) []SpringRow {
    rows := make([]SpringRow, 0)

    scanner := bufio.NewScanner(strings.NewReader(input))
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        line := scanner.Text()

        parts := strings.Split(line, " ")
        row := SpringRow{
            row: parts[0],
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
