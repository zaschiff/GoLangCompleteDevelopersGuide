package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	/*
		var alex person
		alex.firstName = "Alex"
		alex.lastName = "Anderson"
		fmt.Println(alex)
		fmt.Printf("%+v", alex)
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

/*
*************************************************** NOTES ****************************************************
	A Struct is a data structure or a collection of properties that are related together.
	ie: a card could have been made as a struct where the card is a struct
		with properties such as a a suit which may be a type string and a value which
		could also be a string.

		using a struct syntaxt to create a basic stuct of type 'person'
			type   -> keyword declaring a new stype
			person -> name of the new custom type
			struct -> underlying data type
			{}     -> struct properties collection

		properties breakdown: Using the same above example
			firstName -> name of property
			string    -> property type
			lastName  -> name of another property
			string    -> property type

		one way to declare the new struct using th same above example
			alex       -> variable
			person     -> new struct type
			{}         -> property collection, struct body
			"Alex"     -> first property initlaizaiton
			"Anderson" -> second property initializaion

		anotther way
			alex                 -> variable
			person               -> new struct type
			{}                   -> property collection, struct body
			firstName:"Alex"     -> first property initialization explicit to the property name
			LastName:"Anderson"  -> secondd property initialization explicit to the property name

		a third way to assing properties to a struct using the same above example
			var    -> keyword declaring a new variable
			alex   -> variable name
			person -> data stype fo the variable

	to update a property use dot notation ie: alex.firstName
	structs can be embedded in another struct.
	trailing comma needs to be in struct declaration for multiproperty structs regardless
		if it is last value
	field names are not necessary for stuct properties
	functions with receivers can accept a custom struct type

	receivers dont actually pass the value, they passes a copy of the value in a different memory location.
		The copy is then manipulated not the desired value. If we design a receiver with a pointer go
		will allow us to call the function with either the pointer or the value of the pointer type.

	& is an pointer or reference operator. It is used with a varibale name. It gains access to the memory
		address of the varibale specified, ie: &jim is the memory address for the variable jim.

	* is another operator. it is used to dereference a pointer. ie it gets the variable at a memory address
		ie: *(&jim) gives the value of jim. A * in front of a type, means we are looking to a
		pointer to that type, NOT THE VALUE.

	another way to think of pointer operators
		Turn address -> value use *address
		Turn value   -> address use &value

	when a vriable is create but not initialized, go sets it to the zero vaule based on the data type.

	reference types, types that we do not need to worry about pointers
		reference types:
			slices
			maps
			channels
			pointers
			functions

	value types (basic types + structs), we need to use poinrters to manipulate
		value types:
			int
			float
			string
			bool
			structs
*/
