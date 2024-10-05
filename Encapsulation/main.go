package main

import (
	"fmt"

	"github.com/vipinnsingh/go-oops/Encapsulation/structs"
)

/*
	type Person struct {
		firstName string
		lastName  string
		age       int
	}
*/

func main() {

	/*
		p := Person{
			firstName: "",
			lastName:  "",
			age:       0,
		}
	*/

	p := structs.NewPerson("Vipin", "Singh", 23)
	fmt.Printf("p: %v\n", p)

	p.SetFirstName("Updated")

	firstName := p.GetFirstName()
	fmt.Printf("firstName: %v\n", firstName)

}

/*
   if firstName and lastName of Person is not capital then they can not be used/imported in other packages,
   although they can be accessed in same package.
   This is what Encapsulation is.
   So How we can assign or access value of firstName & lastname , is through setter and getter function.

   Summary

   Go supports encapsulation, the OOP principle where data and methods are bundled together, and access to the data is restricted.
   This is achieved through exported (public) and unexported (private) fields and methods.
   In Go, visibility is controlled by capitalization: If a field or method starts with an uppercase letter, it’s exported (public).
   If it starts with a lowercase letter, it’s unexported (private).
*/
