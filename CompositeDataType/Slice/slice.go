package main

import (
	"fmt"
	"strings"
)

func main() {

}

func slice() {

	myArray := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	// mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// mySlice1 := make([]int, 8)

	// mySlice2 := make([]int, 8, 16)

	slc1 := myArray[2:6]

	myArray[2] = 20

	// fmt.Printf("%v\n", myArray)
	// fmt.Printf("%v\n", slc1)

	// fmt.Println("slc1 length: ", len(slc1))
	// fmt.Println("slc1 cap: ", cap(slc1))

	// fmt.Println("myArray length: ", len(myArray))
	// fmt.Println("myArray cap: ", cap(myArray))

	slc1 = append(slc1, 99)
	slc1 = append(slc1, 98)
	slc1 = append(slc1, 97)
	slc1 = append(slc1, 96, 95, 94, 93, 92, 91, 90)

	myArray[7] = 101

	fmt.Printf("%v\n", myArray)
	fmt.Printf("%v\n", slc1)

	fmt.Println("slc1 length: ", len(slc1))
	fmt.Println("slc1 cap: ", cap(slc1))

	fmt.Println("myArray length: ", len(myArray))
	fmt.Println("myArray cap: ", cap(myArray))
}


func funcs() {

	numbers := []int{1, 2, 3, 4, 5}

	addItem(&numbers)

	fmt.Printf("%v \n", numbers)

}

func changeNumbers(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 3
	}
}

func addItem(numbers *[]int){
	*numbers = append(*numbers, 6)
}

func loop() {

	names := []string{"Hamed", "Farzaneh", "Reza", "Sara"}

	for _, item := range names {
		item = strings.ToUpper(item)
	}

	fmt.Printf("%v \n", names)

	for i, item := range names {
		names[i] = strings.ToUpper(item)
	}

	fmt.Printf("%v \n", names)
}

func copySlice() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers2 := []int{10, 20, 30, 40, 50}

	count := copy(numbers, numbers2)

	fmt.Printf("%v \n", numbers)
	fmt.Printf("%d \n", count)

}


func margeSlice() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	numbers = append(numbers, 6, 7, 8, 9)

	fmt.Printf("%v \n", numbers)
	
	numbers = append(numbers, numbers2...)
	
	fmt.Printf("%v \n", numbers)
}


func others() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// 1, 2, 3, 4, 5, 6, 7, 75, 8, 9, 10, 11, 12, 13, 14, 15
	numbers = append(numbers, 0)
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0

	_ = copy(numbers[8:], numbers[7:])
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0
	// 1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 10, 11, 12, 13, 14, 15

	numbers[7] = 75

	fmt.Printf("%v \n", numbers)

	numbers = append(numbers[:7],numbers[8:]...)

	fmt.Printf("%v \n", numbers)


	numbers = numbers[:0]
	fmt.Printf("%v \n", numbers)
	fmt.Printf("%d \n", len(numbers))
	fmt.Printf("%d \n", cap(numbers))

	

	numbers = nil
	fmt.Printf("%v \n", numbers)
	fmt.Printf("%d \n", len(numbers))
	fmt.Printf("%d \n", cap(numbers))


}