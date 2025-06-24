package data

import "givebox/infrastructure/database/profile/user"

var Users = []user.User{
	{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Password:    "hashedPassword",
		PhoneNumber: "081234567890",
		City:        "Jakarta",
		Biography:   "Seorang donatur yang baik hati.",
	},
	{
		FirstName:   "Jane",
		LastName:    "Smith",
		Email:       "jane.smith@example.com",
		Password:    "hashedPassword",
		PhoneNumber: "089876543210",
		City:        "Surabaya",
		Biography:   "Suka berbagi dengan sesama.",
	},
	{
		FirstName:   "Peter",
		LastName:    "Jones",
		Email:       "peter.jones@example.com",
		Password:    "hashedPassword",
		PhoneNumber: "081112233445",
		City:        "Bandung",
		Biography:   "Mari bantu mereka yang membutuhkan.",
	},
}
