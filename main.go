package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// var bookings [50]string //array -> fixed size
// var bookings []string //slice -> like vector
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	for {
		userFirstName, userLastName, userEmail, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userFirstName, userLastName, userEmail, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userFirstName, userLastName, userEmail, userTickets)

			wg.Add(1)
			go sendTicket(userTickets, userFirstName, userLastName, userEmail)
			//prints first name of each user
			printFirstName()

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address entered is wrong")
			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaing, so you can't book %v tickets\n", remainingTickets, userTickets)
			}

		}

	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstName() {
	firstNames := []string{}
	for _, booking := range bookings {
		// var name = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The first name of bookings are : %v\n", firstNames)
}

func validateUserInput(userFirstName string, userLastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&userFirstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Print("Enter the number of tickets you want: ")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTicket(userFirstName string, userLastName string, userEmail string, userTickets uint) {
	remainingTickets -= userTickets
	// bookings[0] = userFirstName + " " + userLastName
	var userData = UserData{
		firstName:       userFirstName,
		lastName:        userLastName,
		email:           userEmail,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is : %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation mail at %v\n", userFirstName, userLastName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}
