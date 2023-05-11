package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type Address struct {
	Firstline string
	Road      string
	City      string
	Pin       string
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, on %s, %s City, @%s", a.Firstline, a.Road, a.City, a.Pin)
}

type Person struct {
	Name    string
	Surname string
	Age     int
	Address *Address
	Job     string
	Senior  bool
}

func (person *Person) IsAddressCityNotSet() bool {
	return person.Address == nil || len(person.Address.City) == 0
}

func (person *Person) HideAddress() {
	if person.Address == nil {
		return
	}
	person.Address.Firstline = "*****"
	person.Address.Road = "*****"
	person.Address.City = "*****"
	person.Address.Pin = "*****"
}

func (person *Person) String() string {
	return fmt.Sprintf("%s %s working as %s at age %d and living at [%s]. Senior:%t.", person.Name, person.Surname, person.Job, person.Age, person.Address, person.Senior)
}

func main() {

	people := make([]*Person, 10)
	names := []string{"nilesh", "rocky", "sam", "tom", "salman", "jack", "jones", "steve"}
	surnames := []string{"peterson", "patil", "longbottom", "khan", "jobs", "dorsey", "musk", "injulkar"}
	jobs := []string{"Engineer", "Politician", "Farmer", "Officer", "CEO", "Electrician", "Watchman", "Plumber", "Doctor"}
	cities := []string{"pune", "mumbai", "london", "delhi", "boston", "seatle"}

	fmt.Printf("Processing %d people %d\n", len(people), cap(people))
	count := cap(people)
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		name := names[r.Intn(len(names))]
		surname := surnames[r.Intn(len(surnames))]
		job := jobs[r.Intn(len(jobs))]
		var city string
		if i%2 == 0 {
			city = cities[r.Intn(len(cities))]
		}
		people[i] = &Person{
			Name:    name,
			Surname: surname,
			Job:     job,
			Age:     i * 10,
		}
		address := &Address{
			Road: "Red Road",
			City: city,
		}
		people[i].Address = address
	}

	fmt.Println("*****BEFORE******")
	for _, person := range people {
		fmt.Println(person)
	}

	// for _, person := range people {
	// 	dataCtx := ast.NewDataContext()
	// 	err := dataCtx.Add("person", person)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	myRules := RuleLib.NewKnowledgeBaseInstance("MyRules", "v1.0")
	// 	Ngn.Execute(dataCtx, myRules)
	// }

	ch := make(chan string)

	for _, person := range people {
		go checkPerson(person, ch)
	}

	for i := 0; i < count; i++ {
		message := <-ch
		fmt.Println("Result from engine " + message)
	}

	fmt.Println("*****AFTER******")
	for _, person := range people {
		fmt.Println(person)
		if person.Age > 60 && !person.Senior {
			panic("Person not processed with senior age " + person.Name)
		}
	}

	for _, person := range people {
		person.Senior = false
		if person.Address != nil {
			person.Address.City = ""
		}
	}

	fmt.Println("*****BEFORE ADDING RULE******")
	for _, person := range people {
		fmt.Println(person)
	}

	AddNewRules()
	RemoveRule("CheckAddress")

	for _, person := range people {
		go checkPerson(person, ch)
	}

	for i := 0; i < count; i++ {
		message := <-ch
		fmt.Println("Result from engine " + message)
	}

	fmt.Println("*****RULE ADDITION AFTER******")
	for _, person := range people {
		fmt.Println(person)
		if person.Age > 60 && !person.Senior {
			panic("Person not processed with senior age " + person.Name)
		}
	}
}

func checkPerson(person *Person, ch chan string) {
	dataCtx := ast.NewDataContext()
	err := dataCtx.Add("person", person)
	if err != nil {
		panic(err)
	}
	myRules := RuleLib.NewKnowledgeBaseInstance("MyRules", "v1.0")
	Ngn.Execute(dataCtx, myRules)
	//Ngn.Execute(dataCtx, MyRules)
	ch <- "Processed " + person.Name
}
