package Services

import (
	"context"
	"database/sql"
	"errors"
	auth "github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Auth"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/paulmach/go.geo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User struct represents a user in the system.
type User struct {
	UserID      int        `json:"user_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	PhoneNumber string     `json:"phone_number"`
	DateOfBirth string     `json:"date_of_birth"`
	Profession  string     `json:"profession"`
	Location    *geo.Point `json:"location"`
	LocDetails  Address    `json:"loc_details"`
	Password    string     `json:"password"`
	ImageId     string     `json:"image_id"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type DBUser struct {
	UserID      int
	FirstName   string
	LastName    string
	PhoneNumber string
	DateOfBirth string
	Profession  string
	Location    *geo.Point
	Password    string
	City        string
	Country     string
	ImageId     string
}

// UserService provides methods to interact with user data.
type UserService struct {
	db *sql.DB
}

// GetAll retrieves all users from the database, including city and country.
func (s *UserService) GetAll(ctx context.Context) ([]User, error) {
	// SQL query to fetch all users, including city and country
	rows, err := s.db.QueryContext(ctx, "SELECT user_id, first_name, last_name, phone_number, date_of_birth, profession, location, city, country, password, profile_image FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	// Iterate over the rows returned from the query
	for rows.Next() {
		var dbUser DBUser
		// Scan the row data into the DBUser struct
		if err := rows.Scan(&dbUser.UserID, &dbUser.FirstName, &dbUser.LastName, &dbUser.PhoneNumber, &dbUser.DateOfBirth, &dbUser.Profession, &dbUser.Location, &dbUser.City, &dbUser.Country, &dbUser.Password, &dbUser.ImageId); err != nil {
			return nil, err
		}

		// Convert DBUser to User, including city and country
		userData := mapDBUserToUser(dbUser)

		// Append the user to the users slice
		users = append(users, userData)
	}

	// Check if any rows were found
	if len(users) == 0 {
		return []User{}, nil // If no users found, return an empty slice
	}

	return users, nil // Return the populated slice of users
}

// GetById retrieves a user by their ID, including city and country.
func (s *UserService) GetById(ctx context.Context, id int) (User, error) {
	// Prepare the query to fetch the user by ID
	query := `SELECT user_id, first_name, last_name, phone_number, date_of_birth, profession, location, city, country, password, profile_image
              FROM users WHERE user_id = ?`

	// Execute the query
	row := s.db.QueryRowContext(ctx, query, id)

	// Initialize the DBUser struct to store the result from the query
	var dbUser DBUser

	// Scan the result into the dbUser struct
	err := row.Scan(&dbUser.UserID, &dbUser.FirstName, &dbUser.LastName, &dbUser.PhoneNumber,
		&dbUser.DateOfBirth, &dbUser.Profession, &dbUser.Location, &dbUser.City, &dbUser.Country, &dbUser.Password, &dbUser.ImageId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return a User with zero values if the user doesn't exist
			return User{}, nil
		}
		// Return any other errors
		return User{}, err
	}

	// Convert the DBUser to the User struct
	user := mapDBUserToUser(dbUser)

	// Return the user and nil (no error)
	return user, nil
}

// GetByName retrieves users by their first or last name, including city and country.
func (s *UserService) GetByName(ctx context.Context, name string) ([]User, error) {
	query := `
        SELECT user_id, first_name, last_name, phone_number, date_of_birth, profession, location, city, country, password, profile_image
        FROM users
        WHERE first_name LIKE ? OR last_name LIKE ?`

	// Use wildcard '%' for partial matching with LIKE
	namePattern := "%" + name + "%"

	rows, err := s.db.QueryContext(ctx, query, namePattern, namePattern)
	if err != nil {
		return nil, err // Return error if the query fails
	}
	defer rows.Close()

	var users []User

	// Iterate over the rows and map each row to a User struct
	for rows.Next() {
		var dbUser DBUser
		err := rows.Scan(&dbUser.UserID, &dbUser.FirstName, &dbUser.LastName, &dbUser.PhoneNumber,
			&dbUser.DateOfBirth, &dbUser.Profession, &dbUser.Location, &dbUser.City, &dbUser.Country, &dbUser.Password, &dbUser.ImageId)
		if err != nil {
			return nil, err // Return error if scanning fails
		}

		// Convert DBUser to User, including city and country
		user := mapDBUserToUser(dbUser)
		users = append(users, user)
	}

	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Create adds a new user to the database, including city and country.
func (s *UserService) Create(ctx context.Context, user *User) error {
	// Perform reverse geocoding to get city and country from the location
	city, country, err := Utils.ReverseGeocode(user.Location.Lat(), user.Location.Lng())
	if err != nil {
		return err // Return if reverse geocoding fails
	}

	// Set the location details in the user struct
	user.LocDetails.City = city
	user.LocDetails.Country = country

	// Hash the user's password before storing it
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Set the hashed password in the user struct
	user.Password = hashedPassword

	// Map the User to DBUser to include city and country
	dbUser := mapUserToDBUser(*user)

	// Prepare the SQL query to insert a new user, including city and country
	query := `
		INSERT INTO users (first_name, last_name, phone_number, date_of_birth, profession, location, city, country, password)
		VALUES (?, ?, ?, ?, ?, ST_GeomFromText(?), ?, ?, ?)
	`

	// Convert the location to WKT format
	locationWKT := user.Location.ToWKT()

	// Execute the query
	_, err = s.db.ExecContext(ctx, query, dbUser.FirstName, dbUser.LastName, dbUser.PhoneNumber, dbUser.DateOfBirth, dbUser.Profession, locationWKT, user.LocDetails.City, user.LocDetails.Country, user.Password)
	if err != nil {
		return err
	}

	var user1 = mapUserToDBUser(*user)

	*user = mapDBUserToUser(user1)

	return nil
}

// Delete removes a user from the database by their ID and returns a boolean to indicate if a user was deleted.
func (s *UserService) Delete(ctx context.Context, userID int) (bool, error) {
	// SQL query to delete a user by ID
	query := `DELETE FROM users WHERE user_id = ?`

	// Execute the query and get the result
	result, err := s.db.ExecContext(ctx, query, userID)
	if err != nil {
		return false, err // Return any error that occurs
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	// Return true if a row was deleted, false otherwise
	return rowsAffected > 0, nil
}

// Converts DBUser to User, adding location details and hiding the password.
func mapDBUserToUser(dbUser DBUser) User {
	return User{
		UserID:      dbUser.UserID,
		FirstName:   dbUser.FirstName,
		LastName:    dbUser.LastName,
		PhoneNumber: dbUser.PhoneNumber,
		DateOfBirth: dbUser.DateOfBirth,
		Profession:  dbUser.Profession,
		Location:    nil,
		LocDetails: Address{
			City:    dbUser.City,
			Country: dbUser.Country,
		},
		Password: "", // Password should not be exposed when mapping to User
		ImageId:  dbUser.ImageId,
	}
}

// Converts User to DBUser for database storage.
func mapUserToDBUser(user User) DBUser {
	return DBUser{
		UserID:      user.UserID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Profession:  user.Profession,
		Location:    user.Location,
		City:        user.LocDetails.City,
		Country:     user.LocDetails.Country,
		ImageId:     user.ImageId,
	}
}

// Update modifies an existing user's information in the database.
func (s *UserService) Update(ctx context.Context, user *User) error {
	// Reverse geocode the new location to get city and country
	city, country, err := Utils.ReverseGeocode(user.Location.Lat(), user.Location.Lng())
	if err != nil {
		return err // Return the error if geocoding fails
	}

	// Set the retrieved city and country in the user struct
	user.LocDetails.City = city
	user.LocDetails.Country = country

	// Hash the user's password before storing it
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// Set the hashed password in the user struct
	user.Password = hashedPassword

	// Map the User to DBUser to include city and country
	dbUser := mapUserToDBUser(*user)

	// Prepare the SQL query to update the user's information
	query := `
        UPDATE users
        SET first_name = ?, last_name = ?, phone_number = ?, date_of_birth = ?, profession = ?, location = ST_GeomFromText(?), city = ?, country = ?, password = ?, image_id = ?
        WHERE user_id = ?
    `

	// Convert the location to WKT format
	locationWKT := user.Location.ToWKT()

	// Execute the query
	_, err = s.db.ExecContext(ctx, query, dbUser.FirstName, dbUser.LastName, dbUser.PhoneNumber, dbUser.DateOfBirth, dbUser.Profession, locationWKT, user.LocDetails.City, user.LocDetails.Country, user.Password, dbUser.ImageId, dbUser.UserID)
	if err != nil {
		return err
	}

	return nil
}

var jwtKey = []byte(Env.GetString("JWT_KEY", "")) // Replace with your secret key

type Claims struct {
	UserID      int    `json:"user_id"`      // Add user_id to claims
	PhoneNumber string `json:"phone_number"` // Keep phone number for authentication
	jwt.RegisteredClaims
}

// Auth method verifies phone number and password, and returns a JWT if valid
func (s *UserService) Auth(ctx context.Context, phoneNumber, password string) (string, error) {
	// SQL query to check if the user exists with the given phone number
	query := `SELECT user_id, first_name, last_name, phone_number, password FROM users WHERE phone_number = ?`

	// Prepare the query
	var dbUser DBUser
	err := s.db.QueryRowContext(ctx, query, phoneNumber).Scan(&dbUser.UserID, &dbUser.FirstName, &dbUser.LastName, &dbUser.PhoneNumber, &dbUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return "", errors.New("user not found")
		}
		return "", err
	}

	// Compare the password provided by the user with the hashed password from the database
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password))
	if err != nil {
		// Password does not match
		return "", errors.New("incorrect password")
	}

	// Create JWT Claims with UserID
	claims := &Claims{
		UserID:      dbUser.UserID, // Populate UserID here
		PhoneNumber: phoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expiration (24 hours)
			Issuer:    "MinBya3mili",                                      // Set the issuer (can be your app's name or any string)
		},
	}

	// Create the JWT token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	// Return the JWT token
	return tokenString, nil
}

func (s *UserService) GetByPhoneNumber(ctx context.Context, phoneNumber string) (User, error) {
	query := `SELECT user_id, first_name, last_name, phone_number, date_of_birth, profession, location, city, country, password FROM users WHERE phone_number = ?`
	var dbUser DBUser
	err := s.db.QueryRowContext(ctx, query, phoneNumber).Scan(
		&dbUser.UserID, &dbUser.FirstName, &dbUser.LastName, &dbUser.PhoneNumber,
		&dbUser.DateOfBirth, &dbUser.Profession, &dbUser.Location, &dbUser.City, &dbUser.Country, &dbUser.Password,
	)
	if err != nil {
		return User{}, err
	}
	return mapDBUserToUser(dbUser), nil
}
