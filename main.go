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
	case "1":
		day1(args[1], args[2])
	case "2":
		day2(args[1], args[2])
	case "3":
		day3(args[1], args[2])
	case "4":
		day4(args[1], args[2])
	case "5":
		day5(args[1], args[2])
	case "6":
		day6(args[1], args[2])
	case "7":
		day7(args[1], args[2])
	case "8":
		day8(args[1], args[2])
	case "9":
		day9(args[1], args[2])
	case "10":
		day10(args[1], args[2])
	}

	fmt.Printf("Execution time: %s", time.Since(start))
}
