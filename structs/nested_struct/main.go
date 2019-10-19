package main

import "fmt"

type Person struct {
	FirstName   string
	LastName    string
	Birthday    string
	ContactData *ContactData
}

type ContactData struct {
	Phonenumber string
	EMail       string
	Adress      *Adress
}

type Adress struct {
	City        string
	Street      string
	HouseNumber string
	State       string
}

func main() {
	var addressbook []Person

	addressbook = append(addressbook, Person{
		FirstName: "Anna",
		LastName:  "Random",
		// pointer to check if nil and smaller because of call by reference instead of call by value
		ContactData: &ContactData{
			Phonenumber: "1234567",
			Adress: &Adress{
				State: "NYC",
			},
		},
	})
	addressbook = append(addressbook, Person{
		FirstName: "Bernard",
		Birthday:  "01/01/1999",
		ContactData: &ContactData{
			EMail: "B.Awesome@hotmail",
		},
	})
	addressbook = append(addressbook, Person{
		FirstName: "Mom",
		ContactData: &ContactData{
			Adress: &Adress{
				City:        "Hometowm",
				Street:      "Oldstreet",
				HouseNumber: "1a",
				State:       "NYC",
			},
		},
	})

	for i, person := range addressbook {
		fmt.Printf("%d. %s %s %s\n", i+1, person.FirstName, person.LastName, person.Birthday)
	}

	fmt.Println("\n\nContact Data:")

	for i, person := range addressbook {
		fmt.Printf("%d. %s %s: \n", i+1, person.FirstName, person.LastName)
		if person.ContactData != nil { // because ContactData is stored as a pointer, we can check here if ContactData is set.
			fmt.Printf("\tContactData:\n")
			if person.ContactData.Phonenumber != "" {
				fmt.Printf("\tPhoneNumber: %s \n", person.ContactData.Phonenumber)
			}
			if person.ContactData.EMail != "" {
				fmt.Printf("\tEmail: %s \n", person.ContactData.EMail)
			}
			if person.ContactData.Adress != nil {
				fmt.Printf("\t\tAdress:  \n")
				if person.ContactData.Adress.Street != "" {
					fmt.Printf("\t\tStreet: %s %s \n", person.ContactData.Adress.Street, person.ContactData.Adress.HouseNumber)
				}
				if person.ContactData.Adress.City != "" {
					fmt.Printf("\t\tCity: %s \n", person.ContactData.Adress.City)
				}
				if person.ContactData.Adress.State != "" {
					fmt.Printf("\t\tState: %s \n", person.ContactData.Adress.State)
				}
			}
		}

	}
}
