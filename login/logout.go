package login

import (
	"bufio"
	"fmt"
	"golang4/auth"
	"os"
)

func Logout(currentUser *auth.User) {
	currentUser = nil
	fmt.Println("logOut success, press enter to back...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
