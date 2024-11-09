package Services

import "database/sql"

type User struct {
	UserID      int    `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"` // Format as "2006-01-02" when converting dates
	Profession  string `json:"profession"`
	Location    Point  `json:"location"` // Can use custom parsing for POINT if needed
}

type UserService struct {
	db *sql.DB
}
