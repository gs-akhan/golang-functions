package main

import "fmt"


//Struct to hold person details

type Person struct {
	name string
	lastname string
}
//to greet people
func Greeting (name string) {
	fmt.Println(name)
}

//function with return value
func PrefixWithsalutation(name string, pre string) (string){
	return name + " " +  pre
}

//function with multiple return values

func GetBothNames(person Person) (string, string){

	return person.name, person.lastname
}


//functions that takes function as argument

func GreetViaFunction(name string, do func(string)) {

	do(name)
}

//function with named return values

func NamedReturnValues(name string, lname string) (a string, b string) {
	a = name
	b = lname

	return
}


//spread operator function 

func SpreadPrint(names ...string) {
	fmt.Println(names[0], names[1])
}

func main() {
	
	Greeting("Azhar")

	Greeting(PrefixWithsalutation("Holla ", "!!!"))

	azhar := Person {"Azharuddin", "Khan"}
	
	fName, lName := GetBothNames(azhar)

	Greeting(fName)
	Greeting(lName)

	GreetViaFunction("Printing Via function", Greeting)

	aName, bName := NamedReturnValues("John", "Ceena")
	
	SpreadPrint(aName, bName)

}