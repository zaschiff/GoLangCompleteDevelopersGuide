package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}

	fmt.Println(link, "is up!")
	ch <- link
}

/*
*************************************************** NOTES ****************************************************
	GO Routines
	When we launch a go program we automatically create a go routine. Its a process that runs out code.
	To start a new go routine we use the keyword go before the code we want it to run.

	Concurrency - load multiple routines at a time even if on a single CPU, executing when possible
	Parallelism - multiple things at the same time executing at the same time.

	child routines are not given the same respect as the main program. if the main program finished before
		the child routines finihsed the childs are killed with out finishing

	DON'T ACCESS THE SAME VARIABL! NEVER SHARE, USE PASS BY VAULE!

	channels are used to communicate between the different go routines. They are the only to communicate
		between the routines.
	channels are like anything else suchs as structs or slices. They are like variables and must can
		only share the type they are designed for.
	if a channle has nothing in it it is deemed complete
	we need a channel for each go routine we have in place

	Syntax for communicating with a channel:
		channel <- 5            ---> this sends the value '5' into the channel
		myNumber <- channel     ---> wait for a number to appear in the channel. when there is a number,
			assign it ot he variable myNumber
		fmt.Println(<- channel) ---> wait for a value to be sent into the channel.,
			when there is one log it immediately

	a for loop with no arguments is infinite

	Function Literal:
		an unnamed function that can be run like a function.
		syntax using the following code func () () {} ()
			func  -> keyword for a function call
			()    -> arugment list
			()    -> return list
			{}    -> function body
			()    -> necessary parentheses to run the function literal

*/
