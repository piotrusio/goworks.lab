package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

type Employee struct {
	Person
	ID int
}

func (p Person) Greet() {
    fmt.Println("Hello, I'm " + p.Name)
}

func (e Employee) Greet() {
    fmt.Println("Hello, I'm Employee " + e.Person.Name)
}



func main() {
	emp := Employee{
		Person: Person{Name: "Piotr", Age: 41},
		ID: 84092708592,
	}
	emp.Person.Greet()
}