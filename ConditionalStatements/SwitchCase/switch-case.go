package main

import (
	"fmt"
	"strings"
)

func main() {
	var score float64

	println("Enter your score: ")
	fmt.Scanln(&score)

	switch {
	case score >= 16 && score <= 20:
		println("A")
	case score >= 11 && score <= 15.99:
		println("B")
	case score >= 6 && score <= 10.99:
		println("C")
	case score >= 0 && score <= 5.99:
		println("D")
	default:
		println("Unknown")
	}
}

func provinceId() {

	var num int
	var provinceName string

	println("Please enter num: ")

	fmt.Scanln(&num)

	switch num {
	case 72, 82, 92:
		provinceName = "Mazandaran"
	case 11, 22, 33, 44, 55, 66, 77, 88, 99, 10, 20, 30, 40, 50, 60, 70, 80, 90:
		provinceName = "Tehran"
	case 13, 23, 43, 53:
		provinceName = "Isfahan"
	case 47, 57, 67:
		provinceName = "Markazi"
	default:
		provinceName = "Unknown"
	}

	println(provinceName)
}

func salaryBySwitchCase() {

	var salary float64
	var minSalary float64 = 5_600_000
	var taxPercent float64 = 0

	fmt.Print("Enter your salary: ")
	fmt.Scanln(&salary)

	switch {
	case salary <= minSalary:
		taxPercent = 0
	case salary > minSalary && salary <= minSalary*2:
		taxPercent = 0.05
	case salary > minSalary*2 && salary <= minSalary*3:
		taxPercent = 0.07
	case salary > minSalary*3 && salary <= minSalary*4:
		taxPercent = 0.10
	default:
		taxPercent = 0.15
	}

	fmt.Printf("Your tax percent is: %.2f\n", taxPercent)
	fmt.Printf("Your tax is: %.2f\n", taxPercent*salary)

	fmt.Printf("Your salary is: %.2f\n", salary-taxPercent*salary)

}



func switchByBreak() {
	var score float64

	println("Enter your score: ")
	fmt.Scanln(&score)

	switch {
	case score >= 16 && score <= 20:
		println("A")
		break
	case score >= 11 && score <= 15.99:
		println("B")
	case score >= 6 && score <= 10.99:
		println("C")
	case score >= 0 && score <= 5.99:
		println("D")
	default:
		println("Unknown")
	}
}


func switchByBadFallthrough() {

	var notificationType string // "sms,email,push"

	println("Enter notification type: ")

	fmt.Scanln(&notificationType)

	switch {
	case strings.Contains(notificationType, "sms"):
		println("SMS sent")
		fallthrough
	case strings.Contains(notificationType, "email"):
		println("Email sent")
		fallthrough
	case strings.Contains(notificationType, "push"):
		println("Push sent")
	default:
		println("Unknown")
	}

}


const monthDays1 int = 31
const monthDays2 int = 30
const monthDays3 int = 29

func switchByGoodFallthrough() {

	var month int
	var totalDays int = 0

	println("Please enter a month number: ")

	fmt.Scanln(&month)

	switch month {
	case 12:
		totalDays += monthDays3
		fallthrough
	case 11:
		totalDays += monthDays2
		fallthrough
	case 10:
		totalDays += monthDays2
		fallthrough
	case 9:
		totalDays += monthDays2
		fallthrough
	case 8:
		totalDays += monthDays2
		fallthrough
	case 7:
		totalDays += monthDays2
		fallthrough
	case 6:
		totalDays += monthDays1
		fallthrough
	case 5:
		totalDays += monthDays1
		fallthrough
	case 4:
		totalDays += monthDays1
		fallthrough
	case 3:
		totalDays += monthDays1
		fallthrough
	case 2:
		totalDays += monthDays1
		fallthrough
	case 1:
		totalDays += monthDays1

	}

	println("Total days: ", totalDays)
}