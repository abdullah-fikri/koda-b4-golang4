package login

import (
	"bufio"
	"fmt"
	"golang4/auth"
	"os"
)

func (s *Service) ListUser(data *[]auth.User) {
	fmt.Println("\n--- List all users ---")
	if len(*data) == 0 {
		fmt.Println("No users registered.")
		return
	}
	for i, user := range *data {
		fmt.Printf("%d. Name: %s %s \n email: %s\n Password: %s", i+1, user.FirstName, user.LastName, user.Email, user.Password)
	}
	fmt.Print("\nPress enter to back...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
