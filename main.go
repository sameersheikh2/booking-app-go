package main

import (
	"fmt"
	"strconv"
	"strings"
)

// its called Package level variables because it is available locally for every function in this file
var totalTickets int = 50
var ticketsAvailable uint = 50 //uint var type allow only positive number
// var bookings [50]string        //this is array.first it need array size and second is value type that comes in array
// var bookings []string //slice in go, same as array but dynamic and no need to define size as we do in array
var bookings = make([]map[string]string, 0) //map["key data type"]value data type;=syntax;; here we created list of maps inside list of slice thus []map[string]string, 0, where first [] impies empty slice and map[]string means slice of maps

func main() {

	//assign variable wihtout datatype because go will automatically understand data type by looking at value assign

	fmt.Println("Welcome to Go Conference booking application")

	fmt.Printf("We have total of %v tickets and %v are still available\n", totalTickets, ticketsAvailable) //%v stand for value. formatting using printf

	fmt.Println("Get your tickets here to attend")

	for {
		firstName, lastName, email, userTickets := getInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTickets)
			firstName := firstNames()
			fmt.Printf("The first name of bookings are %v\n", firstName)

			if ticketsAvailable == 0 {
				fmt.Println("The tickets are all sold out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}

	}

}

func getInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName, lastName, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= ticketsAvailable
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {

	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["bookedTickets"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)
	ticketsAvailable = ticketsAvailable - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for Go Conference.\n", ticketsAvailable)
	fmt.Printf("users are:%v\n", userData)

}

func firstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// OLD CODE:- var names = strings.Fields(booking)   //string.Fields split strings eg. sameer sheikh = "sameer" "sheikh"
		firstNames = append(firstNames, booking["firstName"]) //booking holding value eg. [{firsName:"sameer" ,lastName:"sheikh",email:"sameersheikh@gmail.com"}] "sheikh" so name[0] at index 0 we have value sameer
	}
	return firstNames

	// above we defined an slice of strings(firstNames) and looped through BOOKINGS for both its index and value using the for..range keyword

	/* FOR LOOP:-  for initialization; condition; update {
				statement(s)
			  }.
	The initialization initializes and/or declares variables and is executed only once.
	Then, the condition is evaluated. If the condition is true, the body of the for loop is executed.
	The update updates the value of initialization.
	The condition is evaluated again. The process continues until the condition is false.
	If the condition is false, the for loop ends*/
}
