package main

import (
    "fmt"
    "os"
)

func day13(part string, filename string) {
    content, err := os.ReadFile(fmt.Sprintf("data/day13/%s.txt", filename))
    if err != nil {
        panic(err)
    }

    input := string(content)

    fmt.Printf("Input: %s\n", input)

    switch part {
    case "1":
        panic("Not implemented yet")
    case "2":
        panic("Not implemented yet")
    }
}
