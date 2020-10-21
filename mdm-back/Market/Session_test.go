package market

import "testing"

func TestAddUser(t *testing.T) {
	admin := &User{}
	sess := MakeNewSession(admin)
	newUser := &User{Name: "testUser"}
	sess.AddUser(newUser)
	for _, user := range sess.Users {
		if user.Name == newUser.Name {

		}
	}
}
