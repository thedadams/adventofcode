package main

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/thedadams/adventofcode/2023/util"
)

//go:embed input.txt
var f embed.FS

func main() {
	partOne()
	partTwo()
}

func partOne() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	plays := make([]hand, 0)
	for s.Scan() {
		bid, _ := strconv.Atoi(strings.Split(s.Text(), " ")[1])
		h := newHand(strings.Split(strings.Split(s.Text(), " ")[0], ""), bid)
		plays = append(plays, h)
	}

	sort.Sort(hands(plays))
	fmt.Printf("Answer Day Seven, Part One: %v\n", hands(plays).score())
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	plays := make([]hand, 0)
	for s.Scan() {
		bid, _ := strconv.Atoi(strings.Split(s.Text(), " ")[1])
		h := newJokerHand(strings.Split(strings.Split(s.Text(), " ")[0], ""), bid)
		plays = append(plays, h)
	}

	sort.Sort(hands(plays))
	fmt.Printf("Answer Day Seven, Part Two: %v\n", hands(plays).score())
}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type hand struct {
	handType handType
	cards    []int
	bid      int
}

func (h hand) less(other hand) bool {
	if h.handType != other.handType {
		return h.handType < other.handType
	}

	for i := 0; i < len(h.cards); i++ {
		if h.cards[i] != other.cards[i] {
			return h.cards[i] < other.cards[i]
		}
	}

	return false
}

type hands []hand

func (h hands) Len() int {
	return len(h)
}

func (h hands) Less(i, j int) bool {
	return h[i].less(h[j])
}

func (h hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hands) score() int {
	var score int
	for i, t := range h {
		score += (i + 1) * t.bid
	}

	return score
}

func newHand(cards []string, bid int) hand {
	crds := make([]int, len(cards))
	cardCounts := make(map[string]int, len(cards))
	var mode int
	for i, c := range cards {
		crds[i] = cardValues[c]
		cardCounts[c]++
		if cardCounts[c] > mode {
			mode = cardCounts[c]
		}
	}
	h := hand{
		cards: crds,
		bid:   bid,
	}

	switch mode {
	case 5:
		h.handType = fiveOfKind
	case 4:
		h.handType = fourOfKind
	case 3:
		if len(cardCounts) == 2 {
			h.handType = fullHouse
		} else {
			h.handType = threeOfKind
		}
	case 2:
		if len(cardCounts) == 3 {
			h.handType = twoPair
		} else {
			h.handType = onePair
		}
	default:
		h.handType = highCard
	}

	return h
}

func newJokerHand(cards []string, bid int) hand {
	crds := make([]int, len(cards))
	cardCounts := make(map[string]int, len(cards))
	var (
		mode, jokerCount int
	)
	for i, c := range cards {
		if c == "J" {
			// Jokers are worth 1.
			crds[i] = 1
			jokerCount++
		} else {
			crds[i] = cardValues[c]
			cardCounts[c]++
			if cardCounts[c] > mode {
				mode = cardCounts[c]
			}
		}
	}
	h := hand{
		cards: crds,
		bid:   bid,
	}

	switch mode {
	case 0, 5:
		h.handType = fiveOfKind
	case 4:
		h.handType = fourOfKind + handType(jokerCount)
	case 3:
		if len(cardCounts) == 2 {
			// Either 3 of one, 2 of another and no joker, or 3 of one and 1 of another and 1 joker, the latter being four of a kind.
			h.handType = fullHouse + handType(jokerCount)
		} else if jokerCount == 0 {
			// No jokers mean at least 3 different types of cards, so three of a kind.
			h.handType = threeOfKind
		} else {
			// Four or five of a kind, depending on how many jokers there are.
			h.handType = fullHouse + handType(jokerCount)
		}
	case 2:
		if len(cardCounts) == 1 {
			// Three jokers, so five of a kind.
			h.handType = fiveOfKind
		} else if len(cardCounts) == 2 {
			// Two of one kind, two of another, and a joker == full house.
			// Two of one kind, one of another, and two jokers == four of a kind.
			h.handType = threeOfKind + handType(jokerCount)
		} else if len(cardCounts) == 3 {
			// Two, two, one and no jokers == two pair.
			// Two, one, one and one joker == three of a kind.
			h.handType = twoPair + handType(jokerCount)
		} else {
			// Two of one kind and the other three are different == one pair.
			h.handType = onePair
		}
	default:
		// All non-joker cards are different, so the number of jokers determines the number of a kind.
		switch jokerCount {
		case 0:
			h.handType = highCard
		case 1:
			h.handType = onePair
		case 2:
			h.handType = threeOfKind
		case 3:
			h.handType = fourOfKind
		default:
			h.handType = fiveOfKind
		}
	}

	return h
}
