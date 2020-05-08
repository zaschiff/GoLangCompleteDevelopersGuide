package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

/*
*************************************************** NOTES ****************************************************
	if a receiver's variable is nto used, the variable can be omitted. ie eb englishBot -> englishBot

	interfaces: (using the above code)

		type englishBot struct is a type that has a function with a receiver called getGreeting
		type spanishBot struct is a type that has a function with a receiver called getGreeting

		we also have a type bot interface. If there is a type in the program that has a fucntion caller
			getGreeting and this function returns a string, then it is an member of the bot interface and
			in being a member gains acces to the function printGreeting.

	interfaces are used to define a method set. Any type that has this method set is now also the same type
		as the interface.

	Interface breakdown: using the code type bot interface {getGreeting(string, int) (string, error)}
		type            -> keyword declaring  a new custom type
		bot             -> name of the custom  type
		interface       -> keyword declaring this new custom type is an interface
		{}              -> braces to enclose the method set collection
		getGreeting     -> name of function in the method set
		(string, int)   -> arguments the function is required to need in it's call
		(string, error) -> values the function must return to be incuded in the interface

	2 different types that we will start to use:
		concrete types  -> a type that can create a value out of directly
		interface types -> can not directly create a vleu out of this type.

	interfaces are not generic, go does not have a generic type
	interfaces are implicit
	interfaces are a contract to help us manage types ie. GARBAGE IN -> GARBAGE OUT
	interfaces are tough. Step #1 is understanding how to read them
*/
