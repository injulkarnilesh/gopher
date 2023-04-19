package main

type person struct {
	firstName string
	lastName  string
}

func (p person) updateFirstName(newName string) {
	p.firstName = newName
}

func (p *person) updateLastName(newName string) {
	(*p).lastName = newName
}
