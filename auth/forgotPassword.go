package auth

import (
	"fmt"
	"golang4/md5"
)

func (u *User) ForgotPassword(data *[]User) {
	fmt.Println("input your email: ")
	fmt.Scan(&u.Email)

	for i, user := range *data {
		if user.Email == u.Email {
			changePassword(i, *data)
		}
	}
}

func changePassword(i int, data []User) {
	var pass1, pass2 string
	fmt.Print("input new password: ")
	fmt.Scan(&pass1)
	fmt.Print("input confirm password: ")
	fmt.Scan(&pass2)

	if pass1 == pass2 {
		data[i].Password = md5.HashMD5(pass1)
	} else {
		fmt.Println("password not match")
		changePassword(i, data)
	}
}
