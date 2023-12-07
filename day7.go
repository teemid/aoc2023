package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Hand struct {
	cardsShort int
	handType   int
	bid        int
}

func day7(part string, filename string) {
	content, err := os.ReadFile(fmt.Sprintf("data/day7/%s.txt", filename))
	if err != nil {
		panic(err)
	}

	input := string(content)

	switch part {
	case "1":
		hands := day7ParseInput(part, input)
		day7Part1(hands)
	case "2":
		hands := day7ParseInput(part, input)
		day7Part1(hands)
	}
}

func day7Part1(hands []Hand) {
	sort.Sort(ByRank(hands))

	winnings := 0
	for i, hand := range hands {
		winnings += ((i + 1) * hand.bid)
	}

	fmt.Printf("winnings: %d\n", winnings)
}

type ByRank []Hand

func (h ByRank) Len() int      { return len(h) }
func (h ByRank) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByRank) Less(i, j int) bool {
	return h[i].cardsShort < h[j].cardsShort
}

func day7ParseInput(part string, input string) []Hand {
	hands := make([]Hand, 0)

	jokerValue := 0
	switch part {
	case "1":
		jokerValue = 11
	case "2":
		jokerValue = 1
	}

	parseBid := false
	numStart := -1
	numEnd := 0

	hand := Hand{}

	count := [14]int{}

	h := 0

	hIndex := 0
	for idx, char := range input {
		switch char {
		case 'A':
			h = assignCardToHand(h, 14, hIndex)
			hIndex++
			count[13]++
		case 'K':
			h = assignCardToHand(h, 13, hIndex)
			hIndex++
			count[12]++
		case 'Q':
			h = assignCardToHand(h, 12, hIndex)
			hIndex++
			count[11]++
		case 'J':
			h = assignCardToHand(h, jokerValue, hIndex)
			hIndex++
			count[jokerValue-1]++
		case 'T':
			h = assignCardToHand(h, 10, hIndex)
			count[9]++
			hIndex++
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if parseBid {
				if numStart == -1 {
					numStart = idx
					numEnd = idx
				} else {
					numEnd = idx
				}
			} else {
				value := int(char - '0')
				count[value-1]++
				h = assignCardToHand(h, value, hIndex)
				hIndex++
			}
		case ' ':
			parseBid = true
		case '\r':
			continue
		case '\n':
			if numStart != -1 {
				n, err := strconv.Atoi(input[numStart : numEnd+1])
				if err != nil {
					panic(err)
				}

				handType := getHandType(count)

				h |= (handType << 20)

				hand.handType = handType
				hand.bid = n
				hand.cardsShort = h

				h = 0
				hIndex = 0
				numStart = -1
				numEnd = 0
				parseBid = false
				count = [14]int{}
			}

			hands = append(hands, hand)
			hand = Hand{}
		}
	}

	return hands
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

func getHandType(count [14]int) int {
	jokers := count[0]
	cs := count[1:]

	sort.Sort(sort.Reverse(sort.IntSlice(cs)))

	cs[0] += jokers

	switch cs[0] {
	case 5:
		return FIVE_OF_A_KIND
	case 4:
		return FOUR_OF_A_KIND
	case 3:
		if cs[1] == 2 {
			return FULL_HOUSE
		} else {
			return THREE_OF_A_KIND
		}
	case 2:
		if cs[1] == 2 {
			return TWO_PAIR
		} else {
			return ONE_PAIR
		}
	default:
		return HIGH_CARD
	}
}

const bitsPerCard = 4

func assignCardToHand(hand int, value int, index int) int {
	offset := bitsPerCard * (4 - index)
	hand |= (value << offset)

	return hand
}

func printHand(hand *Hand) {
	handType := hand.cardsShort >> 20
	switch handType {
	case FIVE_OF_A_KIND:
		fmt.Printf("Five of a kind")
	case FOUR_OF_A_KIND:
		fmt.Printf("Four of a kind")
	case FULL_HOUSE:
		fmt.Printf("Full house")
	case THREE_OF_A_KIND:
		fmt.Printf("Three of a kind")
	case TWO_PAIR:
		fmt.Printf("Two pair")
	case ONE_PAIR:
		fmt.Printf("One pair")
	case HIGH_CARD:
		fmt.Printf("High card")
	}

	for i := 0; i < 5; i++ {
		v := hand.cardsShort >> ((4 - i) * bitsPerCard)
		v &= 0xF
		fmt.Printf(" ")
		switch v {
		case 14:
			fmt.Printf("A")
		case 13:
			fmt.Printf("K")
		case 12:
			fmt.Printf("Q")
		case 11:
			fmt.Printf("J")
		case 10:
			fmt.Printf("T")
		case 1:
			fmt.Printf("J")
		default:
			fmt.Printf("%d", v)
		}
	}

	fmt.Printf(" (%d)", hand.bid)

	fmt.Printf("\n")
}
