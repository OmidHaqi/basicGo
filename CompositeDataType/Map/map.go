package main

import "fmt"

func main(){


	
}

func creation() {

	// Map Creation:

	names := make(map[string]string)
	names1 := map[string]string{}
	var names2 map[string]string = map[string]string{}

	names["1234567890"] = "Ali Hoessini"
	names1["1234567890"] = "Ali Hoessini"
	names2["1234567890"] = "Ali Hoessini"

	fmt.Printf("%v\n", names)
	fmt.Printf("%v\n", names1)
	fmt.Printf("%v\n", names2)

}


type Person struct {
	Name   string
	Family string
	Age    int
}

func crud() {

	persons := make(map[string]Person)

	persons["1234567890"] = Person{Name: "Ali", Family: "Hoessini", Age: 30}
	persons["4832425354"] = Person{Name: "Reza", Family: "Rezaei", Age: 31}
	persons["8579456213"] = Person{Name: "Milad", Family: "Mohammadi", Age: 31}

	fmt.Printf("%v\n", persons)

	persons["1234567890"] = Person{Name: "Farhad", Family: "Davoodi", Age: 41}

	fmt.Printf("%v\n", persons)

	delete(persons, "8579456213")

	fmt.Printf("%v\n", persons)

	farhad, ok := persons["1234567840"]

	if ok {
		fmt.Printf("%v\n", farhad)
	} else {
		fmt.Printf("Not Found\n")
	}

}


func others() {

	persons := make(map[string]Person)
	personSlice := []string{}
	persons["1234567890"] = Person{Name: "Ali", Family: "Hoessini", Age: 30}
	personSlice = append(personSlice, "1234567890")
	persons["4832425354"] = Person{Name: "Reza", Family: "Rezaei", Age: 31}
	personSlice = append(personSlice, "4832425354")
	persons["8579456213"] = Person{Name: "Milad", Family: "Mohammadi", Age: 31}
	personSlice = append(personSlice, "8579456213")
	persons["9285365246"] = Person{Name: "Peyman", Family: "Rezvani", Age: 31}
	personSlice = append(personSlice, "9285365246")
	persons["6321459957"] = Person{Name: "Amir", Family: "Hamidi", Age: 31}
	personSlice = append(personSlice, "6321459957")
	persons["4563214569"] = Person{Name: "Amin", Family: "Golmahalle", Age: 31}
	personSlice = append(personSlice, "4563214569")

	fmt.Printf("%d\n", len(persons))


	for _, person := range personSlice {
		fmt.Printf("%v\n", persons[person])
	}

	// interview questions:
	// 1. Reference type or value type?
	// 2. Copy map
	persons2 :=  make(map[string]Person)

	for key, value := range persons {
		persons2[key] = value
	}

	// 3. Equality of maps

	// 4. Concurrency:
		// map is not thread safe
		// 1:1000 2:2000 3:3000
		// routine 1: get 1:1000, increase 1500 = 2500, set 2500
		// routine 2: get 1:1000, decrease 500 = 500, set 500 (overwrite)

	// Problem: slice of numbers and target number 
	// sum two numbers


}