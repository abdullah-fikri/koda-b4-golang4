package auth

import (
	"bufio"
	"fmt"
	"golang4/md5"
	"os"
)

type User struct {
	FirstName       string
	LastName        string
	Email           string
	Password        string
	ConfirmPassword string
}

type AuthService interface {
	Register([]User) User
	Login([]User) (*User, bool)
	ForgotPassword(*[]User)
}

func (u *User) Register(data []User) User {
	var confirm, pass, passConfirm string

	fmt.Println("\n--- Register ---")

	fmt.Print("What is your first name: ")
	fmt.Scan(&u.FirstName)

	fmt.Print("What is your last name: ")
	fmt.Scan(&u.LastName)

	fmt.Print("What is your email: ")
	fmt.Scan(&u.Email)

	fmt.Print("Enter a strong password: ")
	fmt.Scan(&pass)

	fmt.Print("Confirm your password: ")
	fmt.Scan(&passConfirm)

	for i := range data {
		if u.Email == data[i].Email {
			fmt.Println("data is used! press enter to back...")
			reader := bufio.NewReader(os.Stdin)
			_, _ = reader.ReadString('\n')
			return u.Register(data)
		}
	}

	if pass == passConfirm {
		u.Password = md5.HashMD5(pass)
		data = append(data, User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Password:  u.Password,
		})
	} else {
		fmt.Println("password not match")
		u.Register(data)
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
