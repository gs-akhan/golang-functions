package main

import "fmt"

func main() {
	var salutationMap map[string]string
	
	salutationMap = make(map[string]string)

	salutationMap['oldname'] = "azharkhan"

	myMap := map[string]string {
		"name" : "azhar",
	}

	myMap["age"] = "12"
	myMap["location"] = "Hyderbad"

	fmt.Println(myMap["location"])
	fmt.Println(salutationMap["oldname"])
}