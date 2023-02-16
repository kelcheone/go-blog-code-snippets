package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	firstname string `required:"true" max:"100"`
	lastname  string
	age       int
}

func (p Person) fullName() string {
	return p.firstname + " " + p.lastname
}

func (p *Person) incrementAge() {
	fmt.Println(p.age) // 30
	p.age++
	fmt.Println(p.age) // 31
}

type Employee struct {
	Person
	Salary int
}

func main() {
	p := new(Person)

	p.firstname = "Kevin"
	p.lastname = "Kelche"
	p.age = 66

	fmt.Println(p.firstname)
	//   fmt.Println((*p).lastname)//Kelche
	fmt.Println((*p).age) // 30
	fmt.Println(p.fullName())

	p.incrementAge()
	fmt.Println(p.age)

	fmt.Println("-----------------------------------------------------------------------")

	emp1 := &Employee{
		Person: Person{
			firstname: "Ruppet",
			lastname:  "Maddof",
			age:       65,
		},
		Salary: 450393,
	}

	fmt.Println(emp1.Person.firstname)

	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("firstname")

	fmt.Println(field.Tag)
}
