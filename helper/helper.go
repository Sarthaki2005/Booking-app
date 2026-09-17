package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remaining uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remaining
	//return any numberofvalues from go function
	return isValidName, isValidEmail, isValidUserTickets
}

//we have to provide both file to go coomand
//go run main.go helper.go
//or instead of listing all dependent files files in one line we can run all the files in folder booking-app

//go run .

//in order that other package should be able to use this package
//u have to export this
//In go just capitalize the function name
