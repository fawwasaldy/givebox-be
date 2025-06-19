package user

import "fmt"

type Name struct {
	FirstName string
	LastName  string
}

func NewName(firstName, lastName string) (Name, error) {
	if firstName == "" || lastName == "" {
		return Name{}, fmt.Errorf("first name and last name cannot be empty")
	}
	return Name{
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}

func NewNameFromSchema(firstName, lastName string) Name {
	return Name{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}
