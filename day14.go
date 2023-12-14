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
    fmt.Println(rocks)
}

func tilt(rocks [][]rune) {
    for y, line := range rocks {
        for x, e := range line {
            fmt.Printf("%d,%d: %c\n", x, y, e)
        }
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

