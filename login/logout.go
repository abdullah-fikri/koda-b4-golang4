package login

import (
	"bufio"
	"fmt"
	"golang4/auth"
	"os"
)

func (s *Service) LogOut(currentUser *auth.User) bool {
	fmt.Println("logOut success, press enter to back...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	return true
}
