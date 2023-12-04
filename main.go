package main

import (
    "fmt"
    "log"
    "os"
    "time"
)

func main() {
    args := os.Args[1:]
    if len(args) != 3 {
        log.Fatal("Usage: go run main.go <day> <part> <input>")
    }

    start := time.Now()

    switch args[0] {
    case "1": day1(args[1], args[2]); break
    case "2": day2(args[1], args[2]); break
    case "3": day3(args[1], args[2]); break
    case "4": day4(args[1], args[2]); break
    }

    fmt.Printf("Execution time: %s", time.Since(start))
}
