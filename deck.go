package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver, any variable with type 'deck' has access to the "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// convert deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",") // join slice of strings to one string

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // 0666 anyone can read/write
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // slice of strings
	return deck(s)                      // convert to deck, s & deck are essentially same

}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] // swap number
	}
	// source := rand.NewSource(time.Now().UnixNano()) // source
	// r := rand.New(source) // seed

	// for i := range d {
	// 	newPosition := r.Intn(len(d) - 1)
	// 	d[i], d[newPosition] = d[newPosition], d[i] // swap number
	// }
}
