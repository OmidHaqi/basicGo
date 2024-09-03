package main

import (
	"fmt"
	"sort"
)

func main() {
    
} 

func anonymous() {

	myFunc := func() {
		fmt.Println("Hello World")
	}

	x := 1200

	myFunc()
	myFunc()

	println(func(numbers ...int) int {
		var total int
		for _, number := range numbers {
			total += number
			x++
			println(&x)
		}
		return total
	}(1, 2, 3, 4, 5))

	println(x)

	sum := func(numbers ...int) int {
		var total int
		for _, number := range numbers {
			total += number
			x++
			println(&x)
		}
		return total
	}

	println(sum(1, 2, 3, 4, 5))
	println(x)
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	println(x)
}


func sortFunc() {

	numbers := []int{12, 5, 48, 32, 1, 2, 74, 36, 90, 100, 95, 42}

	fmt.Printf("%v\n", numbers)

	sortingFunc := func(i, j int) bool {
		return numbers[i] < numbers[j]
	}
	sort.Slice(numbers, sortingFunc)

	fmt.Printf("%v\n", numbers)
}