package main

import "fmt"

func main() {
	i, j := 20, 40

	var ip *int = &i
	var jp *int = &j

	fmt.Println("Address of i:", &i)
	fmt.Println("Address of ip  pointer:", ip)

	fmt.Println("Address of j:", &j)
	fmt.Println("Address of jp pointer:", jp)

	i2 := i
	i2 = i2 + 2
	fmt.Println("value of i after increase i2", i)
	fmt.Println("value of i2 after increase i2", i2)

	fmt.Println("Address of i2:", &i2)

	var ip2 *int = &i

	*ip2 = *ip2 + 2
	fmt.Println("value of ip2", *ip2)
	fmt.Println("value of i", i)

	print("-----------------------------------------------------")	
	
	num := 20
	println("Address of num:", &num)
	changeNumberByValue(num)
	println("value of num:", num)
	
	changeNumberByRef(&num)
	println("value of num:", num)
	
	print("-----------------------------------------------------")
}

func changeNumberByRef(num *int) {
	println("changeNumberByRef Address of num:", num)
	*num = *num + 2
}

func changeNumberByValue(num int) {
	println("changeNumberByValue Address of num:", &num)
	num = num + 2
}
