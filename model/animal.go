package model

type Animal struct {
	UserID      int64  `json:"user_id"`
	AnimalID    int64  `json:"animal_id"`
	Picture     string `json:"picture"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Type        string `json:"type"`
	Breed       string `json:"breed"`
	Shelter     string `json:"shelter"`
	Address     string `json:"address"`
	PostalCode  string `json:"postal_code"`
	Bio         string `json:"bio"`
}
