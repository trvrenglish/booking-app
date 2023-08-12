package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // List of maps with 0 initial size

// Structs let you save mixed data types for any entity (user, conference, etc.) like a lightweight class
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked up. Come back next year!")
		}

	} else {
		fmt.Println("Sorry, your transaction failed because:")
		if !isValidName {
			fmt.Println("	First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("	Email address you entered is in the wrong format")
		}
		if !isValidTicketNumber {
			fmt.Println("	Number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Provide us some info to purchase your tickets below...\n ")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// Get firstName of a booking from bookings with struct with booking.firstName
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user | map[{key type}]{value type}
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("\nThank you, %v %v, for booking %v tickets.\n	You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("	There are %v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(3 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n###############################")
	fmt.Printf("Ticket sent to %v:\n	%v\n", email, ticket)
	fmt.Println("###############################")
	wg.Done()
}
