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

func day4(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day4/%s.txt", filename))
	if err != nil {
		panic(err)
	}

	input := string(content)

	cards := parseDay4Input(input)

	switch part {
	case "1":
		day4Part1(cards)
	case "2":
		day4Part2(cards)
	default:
		panic("Invalid part for day 4!")
	}

}

func day4Part1(cards []Card) {
	points := 0
	for _, card := range cards {
		winningNumberCount := 0
		for _, winningNumber := range card.WinningNumbers {
			if indexOf(card.Numbers, winningNumber) != -1 {
				winningNumberCount++
			}
		}

		if winningNumberCount > 0 {
			points += 1 << (winningNumberCount - 1)
		}
	}

	fmt.Printf("Points: %d\n", points)
}

func day4Part2(cards []Card) {
	totalCards := 0
	for idx, card := range cards {
		winningNumberCount := 0
		for _, winningNumber := range card.WinningNumbers {
			if indexOf(card.Numbers, winningNumber) != -1 {
				winningNumberCount++
				cards[idx+winningNumberCount].Copies += card.Copies
			}
		}

		totalCards += card.Copies
	}

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

func parseDay4Input(input string) []Card {
	cards := make([]Card, 0)

	numberStart := -1
	numberEnd := 0

	card := NewCard()
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
		case '|':
			card.WinningNumbers = current
			numberStart = -1
			current = make([]int, 0)
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
		case '\r':
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
		}
	}

	return cards
}
