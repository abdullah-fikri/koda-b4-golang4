package main

import (
	"fmt"
	"golang4/auth"
	"golang4/login"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprint(r))
			main()
		}
	}()
	data := []auth.User{}
	loginUser := false
	var currentUser *auth.User

	for {
		var input string
		var svc auth.AuthService = &auth.User{}

		if loginUser {

			fmt.Printf("Halo %s", currentUser.FirstName)

			fmt.Print(`
			1. list all users
			2. logout
			
			0. exit
			
			
			Choose a menu: `)
			fmt.Scan(&input)

			switch input {
			case "1":
				login.ListUser(&data)
			case "2":
				login.Logout(currentUser)
				loginUser = false
			case "0":
				os.Exit(0)
			}

		} else {

			fmt.Println(`
				--- Welcome to system---
	
				1. Register
				2. Login
				3. Forgot Password
	
				0. Exit`)
			fmt.Print("Choose a menu:  ")
			fmt.Scan(&input)

			switch input {
			case "1":
				registeredUser := svc.Register(data)
				data = append(data, registeredUser)

			case "2":
				loggedInUser, y := svc.Login(data)
				if y {
					loginUser = true
					currentUser = loggedInUser
				}

			case "3":
				if len(data) == 0 {
					panic("not data yet...")
				}
				svc.ForgotPassword(&data)

			case "0":
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		}
	}
}
