package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Got 6842 steps with input.txt for part 1

type Point struct {
	x int
	y int
}

type PipeMap struct {
	m     [][]rune
	start Point
}

func day10(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day10/%s.txt", filename))
	if err != nil {
		panic(err)
	}

	input := string(content)

	pipeMap := day10ParseInput(input)
	_, s := findConnectedNeighbours(pipeMap.m, pipeMap.start)

	pipeMap.m[pipeMap.start.y][pipeMap.start.x] = s

	fmt.Printf("start is: %c\n", s)

	switch part {
	case "1":
		day10Part1(pipeMap)
	case "2":
		day10Part2(pipeMap)
	}
}

type Pipe struct {
	steps int
	wn    int
}

func day10Part1(pipeMap *PipeMap) {
	steps, _ := search(pipeMap, pipeMap.start)

	fmt.Printf("Steps: %d\n", steps/2)
}

func day10Part2(pipeMap *PipeMap) {
	_, visited := search(pipeMap, pipeMap.start)

	stripped := make([][]rune, len(pipeMap.m))
	for i := range stripped {
		stripped[i] = make([]rune, len(pipeMap.m[i]))
	}

	for p := range visited {
		stripped[p.y][p.x] = pipeMap.m[p.y][p.x]
	}

	wnm := make([][]rune, len(pipeMap.m))
	for i := range wnm {
		wnm[i] = make([]rune, len(pipeMap.m[i]))
	}

	for p, pipe := range visited {
		if pipe.wn == 1 {
			wnm[p.y][p.x] = '↑'
		} else if pipe.wn == -1 {
			wnm[p.y][p.x] = '↓'
		} else {
			wnm[p.y][p.x] = pipeMap.m[p.y][p.x]
		}
	}

	enclosed := 0
	for y, row := range stripped {
		wn := 0

		for x := range row {
			p := Point{x, y}
			pipe, ok := visited[p]
			if ok {
				wn += pipe.wn
				continue
			}

			if !ok && wn == 0 {
				stripped[y][x] = 'O'
			} else if !ok {
				stripped[y][x] = 'I'
				enclosed++
			}
		}
	}

	fmt.Printf("Enclosed: %d\n", enclosed)
}

func printMap(m [][]rune) {
	for _, row := range m {
		printRow(row)
	}
}

func printRow(row []rune) {
	for _, c := range row {
		switch c {
		case 0:
			fmt.Printf(" ")
		default:
			fmt.Printf("%c", c)
		}
	}

	fmt.Printf("\n")
}

func search(pipeMap *PipeMap, start Point) (int, map[Point]Pipe) {
	steps := 1
	visited := make(map[Point]Pipe)

	neighbours, _ := findConnectedNeighbours(pipeMap.m, start)
	visited[start] = Pipe{0, 0}

	stack := []Point{neighbours[1]}
	path := []Point{neighbours[0], start}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		neighbours, _ := findConnectedNeighbours(pipeMap.m, p)
		for _, n := range neighbours {
			if _, ok := visited[n]; !ok {
				stack = append(stack, n)

			}
		}

		visited[p] = Pipe{steps: steps, wn: 0}

		path = append(path, p)

		steps++
	}

	// update winding numbers
	last := path[0]
	for i, p := range path[1:] {
		if c := pipeMap.m[p.y][p.x]; c == '|' || c == 'F' || c == '7' {
			d := diff(last, p)
			v := visited[p]
			if d.y == 0 {
				n := path[(i+2)%len(path)]
				d = diff(p, n)
				v.wn = d.y
			} else {
				v.wn = d.y
			}
			visited[p] = v
		}

		last = p
	}

	return steps, visited
}

func findConnectedNeighbours(m [][]rune, p Point) ([]Point, rune) {
	points := make([]Point, 0)
	maxX := len(m[0])
	maxY := len(m)

	c := m[p.y][p.x]

	var ds []Point = nil
	switch c {
	case '|':
		ds = []Point{{0, -1}, {0, 1}}
	case '-':
		ds = []Point{{-1, 0}, {1, 0}}
	case 'L':
		ds = []Point{{1, 0}, {0, -1}}
	case 'J':
		ds = []Point{{-1, 0}, {0, -1}}
	case '7':
		ds = []Point{{-1, 0}, {0, 1}}
	case 'F':
		ds = []Point{{1, 0}, {0, 1}}
	case 'S':
		ds = []Point{
			{-1, 0},
			{1, 0},
			{0, 1},
			{0, -1},
		}

		dirs := make([]int, 0)
		for i, dir := range ds {
			x := p.x + dir.x
			y := p.y + dir.y
			if x < 0 || x >= maxX || y < 0 || y >= maxY {
				continue
			}

			if m[y][x] == '.' {
				continue
			}

			if dir.y == -1 {
				switch m[y][x] {
				case '|', '7', 'F': // valid
				default:
					continue // invalid
				}
			} else if dir.y == 1 {
				switch m[y][x] {
				case '|', 'L', 'J': // valid
				default:
					continue
				}
			}

			if dir.x == -1 {
				switch m[y][x] {
				case '-', 'L', 'F': // valid
				default:
					continue
				}
			} else if dir.x == 1 {
				switch m[y][x] {
				case '-', '7', 'J': // valid
				default:
					continue
				}
			}

			p := Point{x, y}
			dirs = append(dirs, i)

			points = append(points, p)
		}

		ds = []Point{
			{-1, 0},
			{1, 0},
			{0, 1},
			{0, -1},
		}

		var s rune = 0
		switch {
		case dirs[0] == 0 && dirs[1] == 1:
			s = '-'
		case dirs[0] == 2 && dirs[1] == 3:
			s = '|'
		case dirs[0] == 0 && dirs[1] == 2:
			s = '7'
		case dirs[0] == 0 && dirs[1] == 3:
			s = 'J'
		case dirs[0] == 1 && dirs[1] == 2:
			s = 'F'
		case dirs[0] == 1 && dirs[1] == 3:
			s = 'L'
		}

		return points, s
	default:
		panic("invalid map char: " + string(c))
	}

	for _, dir := range ds {
		x := p.x + dir.x
		y := p.y + dir.y
		if x < 0 || x >= maxX || y < 0 || y >= maxY {
			continue
		}

		if c := m[y][x]; c == '.' {
			continue
		}

		p := Point{x, y}

		points = append(points, p)
	}

	return points, '1'
}

func comp(a Point, b Point) bool {
	return a.x == b.x && a.y == b.y
}

func diff(a Point, b Point) Point {
	return Point{a.x - b.x, a.y - b.y}
}

func day10ParseInput(input string) *PipeMap {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)

	pipeMap := &PipeMap{
		m: make([][]rune, 0),
	}

	start := Point{0, 0}
	for scanner.Scan() {
		line := scanner.Text()

		rl := []rune(line)
		pipeMap.m = append(pipeMap.m, rl)

		for _, c := range line {
			switch c {
			case 'S':
				pipeMap.start = start
			}

			start.x++
		}

		start.x = 0
		start.y++
	}

	return pipeMap
}
