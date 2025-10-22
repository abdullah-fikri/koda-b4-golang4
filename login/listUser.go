package login

import (
	"bufio"
	"fmt"
	"golang4/auth"
	"os"
)

func ListUser(data []auth.User) {
	fmt.Println("\n--- List all users ---")
	if len(data) == 0 {
		fmt.Println("No users registered.")
		return
	}

	for i := range data {
		fullname := data[i].FirstName + data[i].LastName
		fmt.Printf("%d. \nfullname: %s   \nEmail : %s\n ", i+1, fullname, data[i].Email)
	}
	fmt.Print("\n\nPress enter to back...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
