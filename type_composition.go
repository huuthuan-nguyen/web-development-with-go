package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}

type Middle struct {
	middleName string
}

func (p *Person) PrintName() {
	fmt.Println(p.firstName + " " + p.lastName)
}

type Admin struct {
	Person
	Middle
	role string
}

// override the method 
func (a *Admin) PrintName() {
	// parent method
	a.Person.PrintName()
	fmt.Println(a.firstName + " " + a.lastName + " has role: " + a.role)
}

func main() {
	admin := &Admin{
		Person{
			firstName: "Thuan",
			lastName: "Nguyen",
		},
		Middle{
			middleName: "Huu",
		},
		"Admin",
	}

	admin.PrintName()
}