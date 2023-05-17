package model

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Password  string
	Email     string
}

type LoginAttept struct {
	Email    string
	Password string
}
