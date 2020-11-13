package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generate(t *testing.T) {
	want := `package proxy

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
`
	got, _ := generate("./static_proxy.go")
	//	require.Nil(t, err)
	assert.Equal(t, want, got)
}
