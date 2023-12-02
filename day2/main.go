package main 

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

const (
    RED = iota
    GREEN
    BLUE
)

const (
    TOTAL_RED_COUNT = 12
    TOTAL_GREEN_COUNT = 13
    TOTAL_BLUE_COUNT = 14
)

type CubeSet struct {
    Color int
    Count int
}

type Reveal struct {
    CubeSets []CubeSet
}

type Game struct {
    ID int
    Reveals []Reveal
}

func main() {
    args := os.Args[1:]
    if len(args) != 2 {
        fmt.Println("Please provide a file name and part number")
    }

    content, err := os.ReadFile(args[0])
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    games := parseInput(content)

    if args[1] == "1" {
        part1(games)
    } else if args[1] == "2" {
        part2(games)
    } else {
        fmt.Println("Please provide a valid part number")
    }
}

func part1(games []Game) {
    sum := 0
    for _, game := range games {
        if isValidGame(&game) {
            sum += game.ID
        }
    }

    fmt.Println(sum)
}

type MinCubeSet struct {
    red int
    green int
    blue int
}

func part2(games []Game) {
    sum := 0

    for _, game := range games {
        minCubeSet := MinCubeSet{0, 0, 0}
        for _, reveal := range game.Reveals {
            for _, cubeSet := range reveal.CubeSets {
                switch cubeSet.Color {
                case RED: minCubeSet.red = max(minCubeSet.red, cubeSet.Count); break
                case GREEN: minCubeSet.green = max(minCubeSet.green, cubeSet.Count); break
                case BLUE: minCubeSet.blue = max(minCubeSet.blue, cubeSet.Count); break
                }
            }
        }

        power := minCubeSet.red * minCubeSet.green * minCubeSet.blue
        sum += power
    }

    fmt.Printf("Sum: %d\n", sum)
}

func isValidGame(game *Game) bool {
    for _, reveal := range game.Reveals {
        if !isValidReveal(&reveal) {
            return false
        } 
    }

    return true
}

func isValidReveal(reveal *Reveal) bool {
    isValid := true
    for _, cubeSet := range reveal.CubeSets {
        switch cubeSet.Color {
            case RED: isValid = cubeSet.Count <= TOTAL_RED_COUNT; break
            case GREEN: isValid = cubeSet.Count <= TOTAL_GREEN_COUNT; break
            case BLUE: isValid = cubeSet.Count <= TOTAL_BLUE_COUNT; break
        }

        if !isValid {
            break
        }
    }

    return isValid
}

func parseInput(content []byte) []Game {
    games := make([]Game, 0)
    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue
        }

        game := Game{}

        parts := strings.Split(line, ":")
        if len(parts) != 2 {
            fmt.Println("Error: Invalid input: Expected each game to be in the format 'Game <id>: <reveals>'")
            os.Exit(1)
        }
        gameStr := strings.Split(parts[0], " ")

        gameID, err := strconv.Atoi(gameStr[1])
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }

        game.ID = gameID

        revealTrimmed := strings.TrimSpace(parts[1])

        revealStrs := strings.Split(revealTrimmed, ";")
        for _, revealStr := range revealStrs {
            reveal := Reveal{}
            cubeSets := make([]CubeSet, 0)

            cubeSetStrs := strings.Split(revealStr, ",")
            for _, cubeSetStr := range cubeSetStrs {
                cubeSet := CubeSet{}

                parts := strings.Split(strings.TrimSpace(cubeSetStr), " ")
                if len(parts) != 2 {
                    fmt.Println("Error: Invalid input: Expected each cube set to be in the format '<count> <color>'")
                    os.Exit(1)
                }
                    
                count, err := strconv.Atoi(parts[0])
                if err != nil {
                    fmt.Println("Error:", err)
                    os.Exit(1)
                }

                cubeSet.Count = count

                color := parts[1]

                switch color {
                case "red": cubeSet.Color = RED; break
                case "blue": cubeSet.Color = BLUE; break
                case "green": cubeSet.Color = GREEN; break
                }

                cubeSets = append(cubeSets, cubeSet)
            }

            reveal.CubeSets = cubeSets
            game.Reveals = append(game.Reveals, reveal)
        }


        games = append(games, game)
    }

    return games
}
