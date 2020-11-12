package proxy

import "testing"

func TestUserProxy(t *testing.T) {
	up := NewUserProxy(&User{})
	up.Login("user", "pass")
}
