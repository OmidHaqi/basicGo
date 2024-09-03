package main

import "fmt"

func main() {

	var myArr0 [5]int
	var myArr1 [6]int = [6]int{1, 2}
	myArr2 := [7]int{1, 2}
	myArr3 := [...]int{1, 2}

	myArr0[2] = 25
	myArr1[5] = 25
	myArr2[6] = 25
	myArr3[1] = 25

	fmt.Println(myArr0)
	fmt.Println(myArr1)
	fmt.Println(myArr2)
	fmt.Println(myArr3)

	fmt.Println(len(myArr0))
	fmt.Println(len(myArr1))
	fmt.Println(len(myArr2))
	fmt.Println(len(myArr3))

	copyInArray()

}


func copyInArray() {
	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	numbers2 := &numbers

	numbers3 := numbers[:2]

	numbers4 := numbers3

	numbers[0] = 100

	println(&numbers)
	println(&numbers2)
	println(&numbers3)
	println(&numbers4)

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)

	changeValue(&numbers)
	changeValue(numbers2)

	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers2: %v\n", numbers2)

}

func changeValue(mayArray *[8]int) {
	mayArray[0] = 555
	mayArray[1] = 666
}