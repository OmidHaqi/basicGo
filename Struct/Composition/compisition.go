package main

import "fmt"

const (
	BaseSalary       = 5_600_000
	ExtraHourRate    = 90_000
	HourlySalaryRate = 110_000
	ShiftSalaryRate  = 80_000
	RemoteSalaryRate = 50_000
	TaxRate          = 0.09
)

func main() {
	
	fullTimeEmployees := []FullTimeEmployee{
		
		{Employee : Employee{Id: 1, NationalCode: "1234567890", FullName: "Pejman Rezaee",} ,ExtraHours: 40},
		{Employee: Employee{Id: 2, NationalCode: "4836524125", FullName: "Maryam Hosseini"}, ExtraHours: 120},
	}

	partTimeEmployees := []PartTimeEmployee{
		{Employee : Employee{Id: 3, NationalCode: "6563453455", FullName: "Milad Hassani"}, PartTimeHours: 100},
		{Employee : Employee{Id: 4, NationalCode: "5435435435", FullName: "Maryam Rezaee"}, PartTimeHours: 87},
	}

	shiftEmployees := []ShiftEmployee{
		{Employee : Employee{Id: 5, NationalCode: "3123123213", FullName: "Shahin"}, ShiftHours: 150},
		{Employee : Employee{Id: 6, NationalCode: "6363454355", FullName: "Masoud"}, ShiftHours: 168},
	}

	remoteEmployees := []RemoteEmployee{
		{Employee : Employee{Id: 7, NationalCode: "8529637410", FullName: "omid"},RemoteHours: 111},
		{Employee : Employee{Id: 8, NationalCode: "147852369", FullName: "umut"}, RemoteHours: 123},
	}

	var employees []EmployeeSalaryCalculate = []EmployeeSalaryCalculate{}

	for _, employee := range fullTimeEmployees {
		employees = append(employees, employee)
	}
	
	for _, employee := range partTimeEmployees {
		employees = append(employees, employee)
	}
	
	for _, employee := range shiftEmployees {
		employees = append(employees, employee)
	}
	for _, employee := range remoteEmployees {
		employees = append(employees, employee)
	}

	for _, employee := range employees {
		salary, tax := employee.SalaryCalculator()
		fmt.Printf("\nEmployee (%T): %v\nSalary: %f\nTax: %f\n", employee, employee, salary, tax)
	}
 
}

type EmployeeSalaryCalculate interface {
	SalaryCalculator() (salary float64, tax float64)
}

type Employee struct {
	Id           int
	NationalCode string
	FullName     string
}

type FullTimeEmployee struct {
	Employee
	ExtraHours   float64
}

type PartTimeEmployee struct {
	Employee
	PartTimeHours float64
}

type ShiftEmployee struct {
	Employee
	ShiftHours   float64
}
type RemoteEmployee struct {
	Employee
	RemoteHours   float64
}

func (employee FullTimeEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*employee.ExtraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = HourlySalaryRate * employee.PartTimeHours
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee ShiftEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = ShiftSalaryRate * employee.ShiftHours
	tax = salary * TaxRate
	salary += tax
	return
}
func (employee RemoteEmployee) SalaryCalculator() (salary float64, tax float64) {
	salary = ShiftSalaryRate * employee.RemoteHours
	tax = salary * TaxRate
	salary += tax
	return
}