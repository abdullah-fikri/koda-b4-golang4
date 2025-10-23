package login

import "golang4/auth"

type LoginService interface {
	ListUser(*[]auth.User)
	LogOut(*auth.User) bool
}

type Service struct{}
