package session

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// User is the users that make up a game
type User struct {
	Name      string       `json:"username,omitempty"`
	Timestamp *time.Time   `json:"timestamp,omitempty"`
	UUID      *uuid.UUID   `json:"uuid,omitempty"`
	Cookie    *http.Cookie `json:"-"`
}

// NewUser creats and initializes a new user instance
func NewUser(name string) *User {
	currentTime := time.Now()
	playerUUID := uuid.NewV4()
	player := User{Name: name, Timestamp: &currentTime, UUID: &playerUUID}
	return &player
}

// CookieGen creates a Cookie with all of a given User's needed info.
// Used for inital setting or re-setting of a user's session cookie
func (user *User) CookieGen() *http.Cookie {
	return &http.Cookie{
		Name:    "uuid",
		Value:   user.UUID.String(),
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
}
