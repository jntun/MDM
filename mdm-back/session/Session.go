package session

import (
	"errors"
	"fmt"
)

// Session represents the logic that connects all users for a given instance
type Session struct {
	Admin *User   `json:"admin,omitempty"`
	Users []*User `json:"users,omitempty"`
}

// MakeNewSession handles creating a new unique session
func MakeNewSession(admin *User) *Session {
	session := Session{Admin: admin}
	return &session
}

// AddUser appends a user to a Session's Users
func (sess *Session) AddUser(user *User) {
	sess.Users = append(sess.Users, user)
}

// GetUser uses a cookie's uuid string to find the matching
// Users instance UUID and returns that user
func (sess *Session) GetUser(cookieUUID string) (user *User, err error) {
	for _, user := range sess.Users {
		if user.UUID.String() == cookieUUID {
			fmt.Println("User found:", user)
			return user, nil
		}
	}
	return nil, errors.New("Could not match player with UUID " + cookieUUID)
}
