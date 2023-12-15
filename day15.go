package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
    label string
    value int
}

func day15(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day15/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)
    input = strings.Trim(input, "\r\n")
    parts := strings.Split(input, ",")

    switch part {
    case "1":
        sum := 0
        for _, p := range parts {
            h := hash(p)
            fmt.Printf("%s: %d\n", p, h)
            sum += h
        }

        fmt.Printf("Sum: %d\n", sum)
    case "2":
        boxes := make([][]Entry, 256)
        for _, p := range parts  {
            label, op, value := findOperation(p)

            boxIndex := hash(label)
            box := boxes[boxIndex]

            switch op {
            case '-':
                // TODO (Emil): Take lense from box
                for i, entry := range box {
                    if entry.label == label {
                        box = append(box[:i], box[i+1:]...)
                        boxes[boxIndex] = box
                    }
                }
            case '=':
                if len(box) > 0 {
                    found := false
                    for j, entry := range box {
                        if entry.label == label {
                            entry.value = value
                            found = true
                            box[j] = entry
                            boxes[boxIndex] = box
                            break
                        }
                    }

                    if !found {
                        box = append(box, Entry{label, value})
                        boxes[boxIndex] = box
                    }
                } else {
                    box = append(box, Entry{label, value})
                    boxes[boxIndex] = box
                }
            }
        }

        totalFocusingPower := 0
        for i, box := range boxes {
            for j, entry := range box {
                focusingPower := (i + 1) * (j + 1) * entry.value

                fmt.Printf("Focusing power: %d\n", focusingPower)

                totalFocusingPower += focusingPower
            }
        }

        fmt.Printf("Total focusing power: %d\n", totalFocusingPower)
    }
}

func printBoxes(boxes [][]Entry) {
    for i, box := range boxes {
        if len(box) == 0 {
            continue
        }

        fmt.Printf("Box %d: ", i)
        entries := make([]string, 0)
        for _, entry := range box {
            entries = append(entries, fmt.Sprintf("[%s %d]", entry.label, entry.value))
        }

        fmt.Printf("%s\n", strings.Join(entries, " "))
    }

}

func hash(s string) int {
    value := 0
    for _, c := range s {
        value += int(c)
        value *= 17
        value %= 256
    }

    return value
}

func findOperation(line string) (string, rune, int) {
    for i, c := range line {
        switch c {
        case '-':
            label := line[:i]
            op := '-'
            value := -1

            return label, op, value
        case '=':
            label := line[:i]
            op := '='
            value, err := strconv.Atoi(line[i+1:])
            if err != nil {
                panic(err)
            }

            return label, op, value
        default: continue
        }
    }

    panic("No operation found")
}

