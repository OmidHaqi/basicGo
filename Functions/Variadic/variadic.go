package main

import (
	"fmt"
	"reflect"
)

func main() {
}

func reflectFunc(numbers ...interface{}) {
	for _, item := range numbers {

		switch item.(type) {
		case int:
			println("int:", reflect.ValueOf(item).Int())
		case string:
			println("string:", reflect.ValueOf(item).String())
		}
	}
}

func calculator(numbers ...int) (sum int, mul int) {
	mul = 1
	for _, number := range numbers {
		sum += number
		mul *= number
	}
	return
}
func PrintLog(logs ...interface{}) {
	for i, item := range logs {
		fmt.Printf("%d: %v\n", i, item)
	}
}
