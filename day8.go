package main

import (
    "fmt"
    "os"
)

type DesertMapNode struct {
    L string
    R string
}

type DesertMap struct {
    directions string 
    nodes map[string]DesertMapNode
}

func day8(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day8/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)
    desertMap := day8ParseInput(input)

    switch part {
    case "1":
        day8Part1(desertMap)
    case "2":
        day8Part2(desertMap)
    }
}

func day8Part1(m *DesertMap, args... string) int {
    steps := 0
    curr := "AAA"
    goal := "ZZZ"
    if len(args) == 2 {
        curr = args[0]
        goal = args[1]
    } 
   
    l := len(m.directions)

    for {
        c := m.directions[steps % l] 
        n, ok := m.nodes[curr]
        if !ok {
            panic(fmt.Sprintf("Could not find node %s", curr))
        }

        switch c {
        case 'L': curr = n.L
        case 'R': curr = n.R
        }

        if (goal == "ZZZ" && curr == "ZZZ") || (curr[2] == 'Z') {
            fmt.Printf("Found ZZZ in %d steps\n", steps + 1)

            return steps + 1
        }

        steps++
    }
}

func day8Part2(m *DesertMap) { 
    steps := 0
    l := len(m.directions)
    positions := make([]string, 0)

    for key := range m.nodes {
        if key[2] == 'A' {
            positions = append(positions, key, "**Z")
        }
    }

    fmt.Printf("positions: %v\n", positions)

    count := 0
    for key := range m.nodes {
        if key[2] == 'Z' {
            count++
        }
    }

    stepList := make([]int, len(positions))
    for i, pos := range positions {
        fmt.Printf("pos: %s\n", pos)

        steps := day8Part1(m, pos)

        stepList[i] = steps
    }

    fmt.Printf("stepList: %v\n", stepList)

    for {
        c := m.directions[steps % l]

        fmt.Printf("Step: %d\n", steps + 1)
        fmt.Printf("Direction: %c\n", c)
        fmt.Printf("Positions: %v\n", positions)

        // fmt.Printf("positions: %v\n", positions)
    
        count := 0
        isExit := true
        nextPositions := make([]string, len(positions))
        for i, pos := range positions {
            n, ok := m.nodes[pos]
            if !ok {
                panic(fmt.Sprintf("Could not find node %s", pos))
            }

            // fmt.Printf("pos: %s, node: %v\n", pos, n)

            next := ""
            switch c {
                case 'L': next = n.L
                case 'R': next = n.R
            }

            if next[2] != 'Z' {
                isExit = false
            } else {
                count++
            }

            nextPositions[i] = next
        }

        if count == len(positions) {
            panic("All positions are at exit")
        }

        if isExit {
            fmt.Printf("Found exit in %d steps\n", steps + 1)

            return
        }

        positions = nextPositions
        steps++
    }
}

func day8ParseInput(input string) *DesertMap {
    m := &DesertMap{nodes: make(map[string]DesertMapNode)}

    from := ""
    n := DesertMapNode{}

    isFirstLine := true
    isMap := false

    i := 0
    for i < len(input) {
        c := input[i]

        switch c {
        case ' ': i++
        case '(': 
            j, name := parseName(i+1, input)
            if j == -1 {
                panic("Could not parse name")
            }

            n.L = name

            i = j
        case ')': i++
        case '=': i++
        case ',': 
            j, name := parseName(i+2, input)
            if j == -1 {
                panic("Could not parse name")
            }
            
            n.R = name

            i = j
        case '\r': i++
        case '\n':
            if isMap {
                m.nodes[from] = n
                from = ""
                n = DesertMapNode{}
            } 

            if isFirstLine {
                m.directions = from
            }

            if from == "" {
                isMap = true
            }

            i++
            isMap = !isFirstLine
            isFirstLine = false
        default:
            j, name := parseName(i, input)
            if j == -1 {
                panic("Could not parse name")
            }

            if !isMap {
                from = name
            }


            i = j

            if isMap {
                from = name
            }
        }
    }

    return m
}

func parseName(start int, input string) (int, string) {
    for i, char := range input[start:] {
        switch {
        case 'A' <= char && char <= 'Z': continue
        case '1' <= char && char <= '9': continue
        default: 
            name := input[start:start+i]
            return start + i, name
        }
    } 

    return -1, ""
}

