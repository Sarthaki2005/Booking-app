package main

//impoert multiple packages add it on new line
import (
	"booking-app/helper" //go will by default search for package int its modules u have to explicitly tell that this is package created by u  mention path where ur helper lies
	"fmt"
	"strings"
)

// package level variables
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remaining uint = 50
var bookings = []string{}

func main() {

	//store usernames in list
	//bboking list can take up the user values upto 50

	//if u are not assigning value u have to declare the datatype of the variable
	greetUsers()

	for {
		//How to take user input
		firstName, lastName, email, userTickets := getUserInput()
		//User Input Validation
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remaining)
		//For email

		if isValidName && isValidUserTickets && isValidEmail {

			//Arraysbookings[0] = firstName + " " + lastName
			//Slices
			remaining, bookings = bookTickets(firstName, lastName, userTickets, remaining, bookings)
			// firstNames := []string{}

			//it gives index and value back for each element
			//Go ki ek problem hai u have to use declared variables
			//If u want tio have varible but never to use them use underscore__
			// for _, booking := range bookings { //for index,booking:=range of booking
			// 	var names = strings.Fields(booking)

			// 	firstNames = append(firstNames, names[0])
			// }
			firstNames := getFirst()
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
			if !isValidName {
				fmt.Println("Incorrect Name.")
			}
			if !isValidEmail {
				fmt.Println("INcorrcet Email")
			}
			if !isValidUserTickets {
				fmt.Println("Incorrect ticket Number.")
			}
			fmt.Printf("We only have %v  tickets remaining\n", remaining)
			//break   instead of breaking the whole application we just want to skip next steps

		}
	}

	//Arrays are fixed Size static in nature
	//Slices make it dynamic size

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v  tickets and %v are still available\n", conferenceTickets, remaining)
	fmt.Println("Get your tickets here to attend!")

}
func getFirst() []string { //inside parameter block there inputs what is comming and outsize thst there is return type
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])

	}
	//returning values from function
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Loops to repeat The Logic
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTickets(firstName string, lastName string, userTickets uint, remaining uint, bookings []string) (uint, []string) {
	remaining = remaining - userTickets

	//Arraysbookings[0] = firstName + " " + lastName
	//Slices
	bookings = append(bookings, firstName+" "+lastName)
	// firstNames := []string{}
	return remaining, bookings
}

//instead of passing all values to diff functions every time we create a plce to store all var such that they are accessible to all functions  these are called package level variables.

//Go works in packages
//A single package can have multiple files

//as long as the files are in one package they can access each other
//exxample main.go is in booking-app i.e main package

//helper.go is in helpwer package
