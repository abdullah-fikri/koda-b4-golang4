package auth

import (
	"bufio"
	"fmt"
	"golang4/md5"
	"os"
)

func (u *User) Login(data []User) (*User, bool) {
	fmt.Println("--- Login ---")

	fmt.Print("Enter Your Email: ")
	fmt.Scan(&u.Email)

	fmt.Print("Enter Your password: ")
	fmt.Scan(&u.Password)

	found := false
	for i := range data {
		if u.Email == data[i].Email && md5.HashMD5(u.Password) == data[i].Password {
			found = true
			fmt.Println("Login Success, press enter to back...")
			reader := bufio.NewReader(os.Stdin)
			_, _ = reader.ReadString('\n')
			return &data[i], true
		}
	}
	if !found {
		fmt.Println("Wrong Email Or Password, press enter to restart")
		u.Login(data)
	}
	return nil, false
}
