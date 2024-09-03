package main

import (
	"fmt"
	"strings"
)

func main() {

	myPrint()

	myString := "this is golang tutorial go go"

	fmt.Println("EqualFold", strings.Contains(myString, "go1"))
	fmt.Println("EqualFold", strings.ContainsAny(myString, "go1"))
	fmt.Println("EqualFold", strings.Count(myString, "go"))
	fmt.Println(strings.Cut(myString, "go1"))

	myStringArray := strings.Split(myString, " ")

	println("Word Count: ", len(myStringArray))

	for _, item := range myStringArray {
		println(item)
	}

	myStringArray2 := strings.Join(myStringArray, ",")

	println("Join", myStringArray2)

	println("Repeat", strings.Repeat("iran ", 10))

	println("Replace", strings.Replace(myString, "golang", "go", 3))
	println("ReplaceAll", strings.ReplaceAll(myString, "go", "golang"))

	println("Compare", strings.Compare("golang", "golang"))
	println("Compare", strings.Compare("Golang", "golang"))
	println("Compare", strings.Compare("Golang", "GOLANG"))

	println("EqualFold", strings.EqualFold("Golang", "GOLANG"))
	println("EqualFold", strings.EqualFold("Golang", "golang"))

	println("HasPrefix", strings.HasPrefix("Iran", "Ir"))
	println("HasPrefix", strings.HasPrefix("Iran", "IR"))

	println("HasSuffix", strings.HasSuffix("Iran", "an"))
	println("HasSuffix", strings.HasSuffix("Iran", "n"))

	println("Index an", strings.Index("Iran", "an"))
	println("Index r", strings.Index("Iran", "r"))

	println("ToLower", strings.ToLower("IrAn"))
	println("ToUpper", strings.ToUpper("IrAn"))
	println("Title", strings.Title("Iran is a country"))

	println("Trim", strings.Trim("  Iran is a country   ", " "))
	println("TrimLeft", strings.TrimLeft("!!Iran is a country!!", "!"))
	println("TrimRight", strings.TrimRight("!!Iran is a country!!", "!"))
	println("Trim", strings.Trim("!!Iran is a country!!", "!"))

}

func myPrint() {

	name := "Omid"
	age := 28
	nationalCode := 4832145478
	score := 7.5
	print("My name is ", name, " and my age is ", age, " and my national code is ", nationalCode, " and my score is ", score, "\n")
	println("My name is", name, "and my age is", age, "and my national code is", nationalCode, "and my score is", score)
	// println()
	fmt.Printf("My name is %s and my age is %d and my national code is %d and my score is %f\n", name, age, nationalCode, score)

	fmt.Printf("name: My type is %T\n", name)
	fmt.Printf("nationalCode binary is %b", nationalCode)

}