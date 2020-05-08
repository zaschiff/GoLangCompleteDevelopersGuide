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
// which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	// iterate through a slice
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

/*
*************************************************** NOTES ****************************************************
	go allows use to cerate a custom type. These customs types are an extension of the basic types.
	using the code: type deck []string
		type     -> keyword to signal that we are creatign a new 'data type'
		deck     -> name fo the custom 'type'
		[]string -> the actual underlying type that the custom object is made of,
			 in this case a slice of strings

	using code: func (d deck) print() {}
	a receiver is like a method of a class. It gives a type access ot the function.
		func     -> keyword declaring a function
		(d deck) -> receiver, this states that any variable 'd', of type 'deck', can access or call the print
			method
		print()  -> function name
		{}       -> function body
	receivers are generally marked by a one letter or two letter call. it could have been  d or de or dec.
		In this case d is easiest.

		BREAKDOWN OF THE RECEIVER: d is the actual passed variable that the function is working with.
			deck is the type that d is expected to be. in toehr language d would be equivalent
			to the keyword 'this'

	when we have a value but we do not need it such as the index of a loop, go has a special symbol. We can
		use _ to tell the compiler that we know about the value but we just do not need it.

	go has the abiltiy to do type conversion. It takes one type and converts it to the desired type.
	the syntax to use type conversion by writing out the type we want witht eh value we have in
		parentheses ie: []byte("Hi there!") with []byte being the type we want and "Hi there" the value

	The error type is an object the has two values. if everything is ok, then the value will be nil; nil is
		go's variant of null if something went wrong the error object will have something else other than nil.
		Error handling in go must be taken cre of almost immediately and can be tricky. Recommendation: if something
		goes wrong, ask "What should be done?"
			In this case:
				Option 1: Log the error and return a call to newDeck()
				Option 2: Log the error and entirely quit the program.

*/
