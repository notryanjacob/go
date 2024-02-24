package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck' which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, val := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	//return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //deprecated ioutil package
	file, err := os.Create("os_cards") //latest os package
	_, err = file.WriteString(d.toString())
	return err
}

func readFile(filename string) deck {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dataByte), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
