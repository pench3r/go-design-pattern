package proxy

import "fmt"

type IUser interface {
	Login(username string, password string) error
}

type User struct{}

func (u *User) Login(username string, password string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(u *User) *UserProxy {
	return &UserProxy{user: u}
}

func (up *UserProxy) Login(username string, password string) error {
	fmt.Println("check username and password")

	if err := up.user.Login(username, password); err != nil {
		fmt.Println("login failed")
		return nil
	}

	fmt.Printf("login info: %s", username)
	return nil
}
