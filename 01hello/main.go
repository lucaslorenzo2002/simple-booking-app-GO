package main

import (
	"fmt"
	"hello/helper"
	"sync"
	"time"
)

const conferenceName = "go conference"
const totalTickets = 50

var remainingTickets = 50
var bookings = make([]userData, 0)

type userData struct {
	username        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers(conferenceName, totalTickets)

	for {

		username, email, userTickets := getUsersInput()

		validUsername, validEmail := helper.Validation(username, email)

		if !validUsername {
			fmt.Println("username must be at least 2 characthers")
			continue
		}

		if !validEmail {
			println("invalid email")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("we only have %v tickets available, so you cant book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		if userTickets <= 0 {
			fmt.Println("please write a valid number")
			continue
		}

		bookTickets(uint(remainingTickets), uint(userTickets), username, email)

		wg.Add(1)
		go sendingTickets(uint(userTickets), username, email)

		if remainingTickets == 0 {
			fmt.Println("no tickets available")
			break
		}
	}

	wg.Wait()
}

func greetUsers(par1 string, par2 int) {
	fmt.Printf("hello to %v world\n", par1)
	fmt.Printf("we have %v total tickets\n", par2)
	fmt.Println("book your tickets here")
}

func getUsersInput() (string, string, int) {
	var username string
	var email string
	var userTickets uint

	fmt.Println("please write your username")
	fmt.Scan(&username)

	fmt.Println("please write your email")
	fmt.Scan(&email)

	fmt.Println("please write the amount of tickets you want")
	fmt.Scan(&userTickets)

	return username, email, int(userTickets)
}

func bookTickets(par1 uint, par2 uint, par4 string, par5 string) {
	par1 = par1 - par2

	var userData = userData{
		username:        par4,
		email:           par5,
		numberOfTickets: par2,
	}

	bookings = append(bookings, userData)

	fmt.Printf("thank you %v for buying %v tickets, you will recive a validation email to %v\n", par4, par2, par5)
	fmt.Printf("there are %v tickets remaining\n", par1)
	fmt.Printf("these are all of our bookings %v\n", bookings)

}

func sendingTickets(tickets uint, username string, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf(" %v tickets for %v", tickets, username)
	println("################")
	fmt.Printf("sending ticket: \n %v to %v \n", ticket, email)
	println("################")
	wg.Done()
}
