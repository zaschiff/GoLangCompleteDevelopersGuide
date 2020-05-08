package main

func main() {
	/*
		cards := newDeck()

		hand, remainingCards := deal(cards, 5)

		hand.print()
		remainingCards.print()

		cards := newDeck()
		 cards.saveToFile("my_cards")
	*/

	cards := newDeck()
	//cards.shuffle()
	cards.print()
}

/*
*************************************************** NOTES ****************************************************

	variable sytnax explanation:
		using: var card string = "Ace of Spades"

			var             -> keyword declaring that we are about to create a new variable
			card            -> name of the varibale we will be creating
			string 		    -> type of information the variable can hold. In this case, a string
			"Ace of Spades" -> the information or value assigned to the variable.

		Go is a staticly typed language.  A vriable will only ever hold th type declared.
		It will ot change for example form a sring to numbers.
		Go's compiler can be inferred if necessary by the value assigned
			ie var card string = "Ace of Spades" can be written as card := "Ace of Spades"
		:= operator is the short variable delcaration; the operator is only used when creating a new variable.
			It declares and initializes the variable.

	functions that return must state what they return in the function declaration.
	this declaration comes in betweeen () {} in the function declaration and it is the type we want to return
		using: func newCard() string {}
			func      -> keyword to declare that the following code is a function
        	newCard   -> sets the name of the function
			()        -> list of arguments to pass the function
			string    -> type of value to be returned
			{}        -> function body, calling the function runs the code between the curly braces.

	Go has two basic array like structures.
		Array -> A fixed length list of items
		Slice -> An array that can grow or shrink
	both array structures need to be defined with the data type it will hold. Neither can have different
	types in the list ie "cheese", "tomato", "lettuce" is valid. "cheese", 12345, "lettuce" is not valid
	syntax for creating an array:
		[]     -> slice declaration
		string -> type slice will hold
		{}     -> the list of items the slice holds. Can be initialized at declaration or not.

	iterate through the slice using a for loop. Using the code: for i, card := range cards {}
		for 		-> keyword starting the loop
		i 			-> index of the element in the array
		card 		-> current card we're iterating over
		range cards -> take the slice of 'cards' and loop over it
		{} 			-> body of the loop. Code in body will be run each iteration.

	Slices are zero-indexed meaning they start their index at 0 not 1
	go has a helper built into the slice type to selet pieces of the array.
		ie: cards[0:2] is a range of the slice starting from the first index to the second index.
			the first is the number in the brackets is the inclusive starting position, the second number
			is the exclusive ending position. This means, in this case, that we are selecting index 0 and 1.

	most programming languages use OO approach with class definiiotns and method creation.
	Go tries to extend basic data types by adding additional functiosn to it.
	in the bove example project, we want to tell Go to create an array of strings and add a bunch
	of functions specifically made to work with it. A function with a receiver is like a method -
	a fuunction the belongs to an instance of a class.

	the project structure of this project may end up somehting like this:
		'card' folder
			- main.go      -> Code to create and manipulate a deck.
			- deck.go      -> Code that describes what a deck is and how it works.
			- deck_test.go -> code to automatically test the deck.
*/
