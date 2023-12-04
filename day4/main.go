package main

import (
	"fmt"
	"os"
	"strconv"
)

type Card struct {
	ID             int
	Copies         int
	WinningNumbers []int
	Numbers        []int
}

func NewCard() Card {
	return Card{
		Copies:         1,
		WinningNumbers: make([]int, 0),
		Numbers:        make([]int, 0),
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	input := string(content)

	cards := make([]Card, 0)
	card := NewCard()

	numberStart := -1
	numberEnd := 0

	var current []int = nil

	for idx, char := range input {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if numberStart == -1 {
				numberStart = idx
				numberEnd = idx
			} else {
				numberEnd = idx
			}
			break
		case ' ':
			if current != nil && numberStart != -1 {
				n, err := strconv.Atoi(input[numberStart : numberEnd+1])
				if err != nil {
					fmt.Printf("Error converting %s to int\n", input[numberStart:numberEnd+1])
					os.Exit(1)
				}
				current = append(current, n)

				numberStart = -1
			}
			break
		case '|':
			card.WinningNumbers = current
			numberStart = -1
			current = make([]int, 0)
			break
		case ':':
			if numberStart != -1 {
				cardID, err := strconv.Atoi(input[numberStart : numberEnd+1])
				if err != nil {
					fmt.Printf("Error converting %s to int\n", input[numberStart:numberEnd+1])
					os.Exit(1)
				}

				card.ID = cardID
				numberStart = -1
			}

			current = make([]int, 0)
			break
		case '\r':
			break
		case '\n':
			if numberStart != -1 {
				n, err := strconv.Atoi(input[numberStart : numberEnd+1])
				if err != nil {
					fmt.Printf("Error converting %s to int\n", input[numberStart:numberEnd+1])
					os.Exit(1)
				}
				current = append(current, n)
			}

			card.Numbers = current
			numberStart = -1
			current = nil
			cards = append(cards, card)
			card = NewCard()
			break
		}
	}

	totalCards := 0
	points := 0
	for idx, card := range cards {
		winningNumberCount := 0
		for _, winningNumber := range card.WinningNumbers {
			if indexOf(card.Numbers, winningNumber) != -1 {
				winningNumberCount++
				cards[idx+winningNumberCount].Copies += card.Copies
			}
		}

		if winningNumberCount > 0 {
			points += 1 << (winningNumberCount - 1)
		}

		totalCards += card.Copies
	}

	fmt.Printf("Points: %d\n", points)
	fmt.Printf("Total cards: %d\n", totalCards)
}

func indexOf(arr []int, val int) int {
	for idx, item := range arr {
		if item == val {
			return idx
		}
	}

	return -1
}
