package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
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
    gi := day11ParseInput(input)

    switch part {
    case "1":
        day11Part1(gi)
    case "2":
        panic("not implemented")
    }
}

type GalaxyPair struct {
    g1 Point
    g2 Point
}

func day11Part1(gi *GalaxyImage) {
    gi.Expand()
    gi.FindGalaxies()

    sum := 0
    gs := gi.galaxies[:]
    for len(gs) > 1 {
        g1 := gs[0]
        gs = gs[1:]
        for _, g2 := range gs {
            d := distance(g1, g2)
            
            sum += d
        }
    }

    fmt.Printf("sum: %d\n", sum)
}

func printImage(gi *GalaxyImage) {
    for _, row := range gi.image {
        for _, c := range row {
            fmt.Printf("%c", c)
        }
        fmt.Println()
    }
    fmt.Println()
}

func (gi *GalaxyImage) Expand() {
    columns := make([]bool, len(gi.image[0]))
    for c := range columns {
        columns[c] = true
    }

    rows := make([]int, 0)
    for y, row := range gi.image {
        emptyRow := true
        for x, char := range row {
            if char != '.' {
                emptyRow = false
                columns[x] = false
            }
        }

        if emptyRow {
            rows = append(rows, y)
        }
    }

    fmt.Printf("rows: %v\n", rows)

    for i, y := range rows {
        gi.image = slices.Insert(gi.image, y+i, gi.image[y+i])
    }

    fmt.Println("columns: ", columns)

    i := 0
    for x, isEmpty := range columns {
        if isEmpty {
            for y := range gi.image {
                gi.image[y] = slices.Insert(gi.image[y], x+i, '.')

            }

            i++
        } 
    }
}

func (gi *GalaxyImage) FindGalaxies() {
    for y, row := range gi.image {
        for x, char := range row {
            if char == '#' {
                gi.galaxies = append(gi.galaxies, Point{x, y})
            }
        }
    }
}

func distance(p1, p2 Point) int {
    return abs(p1.x - p2.x) + abs(p1.y - p2.y)
}

func day11ParseInput(input string) *GalaxyImage {
    gi := &GalaxyImage{
        galaxies: make([]Point, 0),
        image: make([][]rune, 0),
    }

    scanner := bufio.NewScanner(strings.NewReader(input))
    scanner.Split(bufio.ScanLines)

    y := 0
    for scanner.Scan() {
        line := scanner.Text()
        runes := []rune(line)
        gi.image = append(gi.image, runes)

        y++
    }

    return gi
}
