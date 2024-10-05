package structs

type Person struct {
	firstName string
	lastName  string
	age       int
}

// constructor
func NewPerson(firstName string, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) Walk() string {
	return p.firstName + p.lastName + "is walking."
}
