package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create new type of deck, slice of strings

type deck []string


func newDeck() deck {

	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, suit := range suits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string  {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// byte[] => string => string[] => deck
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle () {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}