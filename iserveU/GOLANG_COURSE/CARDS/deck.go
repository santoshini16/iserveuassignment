package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newdeck() deck {
	cards := deck{}
	cardsuites := []string{"Spades", "Diamonds", "Club", "Hearts"}
	cardvalues := []string{"Ace", "Two", "Three", "Four"}

	for _, suites := range cardsuites {
		for _, value := range cardvalues {
			cards = append(cards, value+" of "+suites)
		}
	}
	return cards
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]

}
func (d deck) tostring() string {
	return strings.Join([]string(d), ",")

}
func (d deck) savetofile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.tostring()), 0666)
}
func newdeckfromfile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {

		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//for swapping
	for i := range d {
		newposition := r.Intn(len(d) - 1)

		d[i], d[newposition] = d[newposition], d[i]

	}

}
func main() {

	cards := newdeck()
	cards.shuffle()
	cards.print()
}
