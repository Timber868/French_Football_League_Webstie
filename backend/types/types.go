package types

import "time"

//This is the type i am going to be using to read json data sent and parse payload
type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID        int       `json:"ID"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     time.Time `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAT"`
}
