package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const LineBreak = "\r\n"

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        os.Exit(1)
    }

    task := "1"
    if len(args) > 1 {
        task = args[1]
    }

    content, err := os.ReadFile(args[0])
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }

    switch task {
    case "1": task1(content); break
    case "2": task2(content); break
    }
}

func task1(content []byte) {
    numbers := make([]int, 0)

    lines := strings.Split(string(content), LineBreak)
    for _, line := range lines {
        if line == "" {
            continue
        }

        digits := make([]string, 0)
        for _, char := range line {
            switch {
            case '0' <= char && char <= '9': 
                digits = append(digits, string(char))
                break
            default: continue
            }
        }

        n := 0
        switch len(digits) {
        case 0: continue
        default: 
            var err error = nil
            n, err = strconv.Atoi(digits[0] + digits[len(digits)-1])
            if err != nil {
                fmt.Printf("Error: %s\n", err)
                os.Exit(1)
            }
            break
        }

        numbers = append(numbers, n)
    }

    sum := 0
    for _, n := range numbers {
        sum += n
    }

    fmt.Printf("Sum: %d\n", sum)
}


type Number struct {
    Value int
    Position int
}

type Digit struct {
    Name string
    Value int
}

func task2(content []byte) {
    digits := []Digit{
        {"one", 1},
        {"two", 2},
        {"three", 3},
        {"four", 4},
        {"five", 5},
        {"six", 6},
        {"seven", 7},
        {"eight", 8},
        {"nine", 9},
        {"0", 0},
        {"1", 1},
        {"2", 2},
        {"3", 3},
        {"4", 4},
        {"5", 5},
        {"6", 6},
        {"7", 7},
        {"8", 8},
        {"9", 9},
    }
   

    ns := make([]int, 0)
    lines := strings.Split(string(content), LineBreak)
    for _, line := range lines {
        numbers := make([]Number, 0)
        if line == "" {
            continue
        }

        for _, digit := range digits {
            index := strings.Index(line, digit.Name)
            if index != -1 {
                numbers = append(numbers, Number{digit.Value, index})
            }

            index = strings.LastIndex(line, digit.Name)
            if index != -1 {
                numbers = append(numbers, Number{digit.Value, index})
            }
        }

        sort.Slice(numbers, func(i, j int) bool {
            return numbers[i].Position < numbers[j].Position
        })

        if len(numbers) == 0 {
            continue
        }

        n := 10 * numbers[0].Value + numbers[len(numbers)-1].Value

        ns = append(ns, n)
    }

    sum := 0
    for _, n := range ns {
        sum += n
    }

    fmt.Printf("Sum: %d\n", sum)
}
