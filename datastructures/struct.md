## Struct Fundamentals

```go
// Struct declaration and creation
type Person struct {
    Name string
    Age  int
    City string
}

// Creation methods
person1 := Person{Name: "Alice", Age: 25, City: "NYC"}    // Named fields
person2 := Person{"Bob", 30, "LA"}                        // Positional (fragile)
person3 := Person{Name: "Carol"}                          // Partial: {Carol 0 ""}
var person4 Person                                        // Zero value: {"" 0 ""}

// Field access and modification
fmt.Println(person1.Name)                                 // "Alice"
person1.Age = 26                                          // Modify field
person1.City = "Boston"

// Nested structs
type Address struct {
    Street, City, Zip string
}

type Employee struct {
    Person  Person                                         // Nested struct
    ID      int
    Address Address
}

emp := Employee{
    Person:  Person{Name: "Dave", Age: 28},
    ID:      12345,
    Address: Address{Street: "123 Main", City: "Boston", Zip: "02101"},
}
fmt.Println(emp.Person.Name)                              // Access nested: "Dave"
fmt.Println(emp.Address.City)                             // Access nested: "Boston"

// Struct methods with different receivers
func (p Person) Greet() string {                          // Value receiver
    return "Hello, I'm " + p.Name
}

func (p *Person) HaveBirthday() {                         // Pointer receiver
    p.Age++                                               // Modifies original
}

person := Person{Name: "Eve", Age: 25}
fmt.Println(person.Greet())                               // "Hello, I'm Eve"
person.HaveBirthday()                                     // person.Age becomes 26

// Struct embedding (composition)
type Student struct {
    Person                                                // Embedded - no field name
    StudentID int
    GPA       float64
}

student := Student{
    Person:    Person{Name: "Frank", Age: 20},
    StudentID: 98765,
    GPA:       3.8,
}
fmt.Println(student.Name)                                 // "Frank" - field promoted
fmt.Println(student.Greet())                              // Method promoted from Person
```