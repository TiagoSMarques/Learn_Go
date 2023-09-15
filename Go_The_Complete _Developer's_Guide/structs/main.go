package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 3892,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// since go is a pass by copy language to change vaule types (int,float,string, str,bool,struct) we need to pass the pointer to the type
// for reference types (slice, map, channels, pointers, funtions) we dont pass the pointer because they are already references to the underlying data
func (p *person) updateName(new_first_name string) {
	p.firstName = new_first_name
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

func ex1() {

	// alex1 := person{"Alex", "Anderson"}
	alex2 := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("%+v\n", alex2)

	var alex3 person
	alex3.firstName = "Alex"
	alex3.lastName = "Anderson"

	fmt.Println(alex2, alex3)
}
