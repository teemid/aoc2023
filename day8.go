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
    part2 := (len(args) == 2)
    if part2 {
        curr = args[0]
        goal = args[1]
        part2 = true
    }

    fmt.Printf("Starting at %s, going to %s\n", curr, goal)
   
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

        if (!part2 && curr == goal) || (part2 && curr[2] == 'Z') {
            fmt.Printf("Found ZZZ in %d steps\n", steps + 1)

            return steps + 1
        }

        steps++
    }
}

func day8Part2(m *DesertMap) {
    positions := make([]string, 0)

    for key := range m.nodes {
        if key[2] == 'A' {
            positions = append(positions, key)
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

        steps := day8Part1(m, pos, "**Z")

        stepList[i] = steps
    }

    answer := lcm(stepList...)

    fmt.Printf("Answer: %d\n", answer)
}

func lcm(args... int) int {
    ans := args[0]
    for i := 1; i < len(args); i++ {
        ans = (args[i] * ans) / gcd(args[i], ans)
    }

    return ans
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }

    return a
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

