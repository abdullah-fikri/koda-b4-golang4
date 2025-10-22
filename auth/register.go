package auth

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	FirstName       string
	LastName        string
	Email           string
	Password        string
	ConfirmPassword string
}

func (u *User) Register(data []User) User {
	var confirm string

	fmt.Println("\n--- Register ---")

	fmt.Print("What is your first name: ")
	fmt.Scan(&u.FirstName)

	fmt.Print("What is your last name: ")
	fmt.Scan(&u.LastName)

	fmt.Print("What is your email: ")
	fmt.Scan(&u.Email)

	fmt.Print("Enter a strong password: ")
	fmt.Scan(&u.Password)

	fmt.Print("Confirm your password: ")
	fmt.Scan(&u.ConfirmPassword)

	for i := range data {
		if u.Email == data[i].Email {
			fmt.Println("data is used! press enter to back...")
		}
	}

	fmt.Printf("\nIs it true?\nFirst name: %s\nLast name: %s\nEmail: %s\n",
		u.FirstName, u.LastName, u.Email)

	fmt.Print("Continue (y/n): ")
	fmt.Scan(&confirm)

	switch confirm {
	case "y":
		fmt.Println("Register Success, press enter to back...")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')
		return *u
	case "n":
		return u.Register(data)
	default:
		fmt.Println("Invalid input, returning to menu...")
		return *u
	}
}
