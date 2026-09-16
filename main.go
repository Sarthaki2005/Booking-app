package main

//impoert multiple packages add it on new line
import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remaining uint = 50

	//store usernames in list
	var bookings = []string{} //bboking list can take up the user values upto 50

	//if u are not assigning value u have to declare the datatype of the variable
	greetUsers(conferenceName, conferenceTickets, remaining)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Loops to repeat The Logic
	for {
		//How to take user input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		//User Input Validation
		isValidInput := ValidateUserInput(firstName, lastName, email, userTickets, remaining)
		//For email

		if isValidInput {

			remaining = remaining - userTickets

			//Arraysbookings[0] = firstName + " " + lastName
			//Slices
			bookings = append(bookings, firstName+" "+lastName)
			// firstNames := []string{}

			//it gives index and value back for each element
			//Go ki ek problem hai u have to use declared variables
			//If u want tio have varible but never to use them use underscore__
			// for _, booking := range bookings { //for index,booking:=range of booking
			// 	var names = strings.Fields(booking)

			// 	firstNames = append(firstNames, names[0])
			// }
			firstNames := getFirst(bookings)
			fmt.Printf("The firts name of bookings %v\n", firstNames)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array Type: %T\n", bookings)
			fmt.Printf("Array Length: %v\n", len(bookings))

			fmt.Printf("Thank You %v %v for booking %v tickets ! You will receive an email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are reamining\n", remaining)

			if remaining == 0 {
				//end Program
				fmt.Printf("Our conference is sold out.\n")
				break //break the indefinite loop
			}
		} else {
			fmt.Printf("We only have %v  tickets remaining\n", remaining)
			//break   instead of breaking the whole application we just want to skip next steps

		}
	}

	//Arrays are fixed Size static in nature
	//Slices make it dynamic size

}

func greetUsers(conferenceName string, conferenceTickets int, remaining uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v  tickets and %v are still available\n", conferenceTickets, remaining)
	fmt.Println("Get your tickets here to attend!")

}
func getFirst(bookings []string) []string { //inside parameter block there inputs what is comming and outsize thst there is return type
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	//returning values from function
	return firstNames
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remaining uint) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remaining
	//return any numberofvalues from go function
	return isValidName && isValidEmail && isValidUserTickets
}
