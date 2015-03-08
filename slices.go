package main

import "fmt"

func main() {
	students := []string {"azhar", "PR", "john", "jenny"}
	

	//Adding more elements to slices using append

	students = append(students, "jennifer")

	printAll(students...)

	//create new slice from another slice

	smallStudents := students[:2]

	fmt.Println("-------------------------")

	printAllLoop(smallStudents)

}


func printAll(allNames ...string) {
	
	for index, value := range allNames {
		if (index %2 == 0) {
			fmt.Println(value)	
		}
	}
	
}

func printAllLoop(allNames []string) {
	
	for i := 0 ; i < len(allNames); i++ {
		fmt.Println(allNames[i])
	}
}

